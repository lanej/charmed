"""Collect serial, real-PTY five-second dashboard measurements."""
import argparse
import errno
import fcntl
import hashlib
import json
import math
import statistics
import os
from pathlib import Path
import pty
import re
import select
import struct
import subprocess
import sys
import termios
import time

ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / 'scripts'))
import perf_common


def seconds(value):
    matched = re.fullmatch(r'(\d+(?:\.\d+)?)(ms|s|m|h)', value)
    if not matched:
        raise argparse.ArgumentTypeError('use a positive duration such as 5s or 2m')
    result = float(matched[1]) * {'ms': .001, 's': 1, 'm': 60, 'h': 3600}[matched[2]]
    if result <= 0:
        raise argparse.ArgumentTypeError('duration must be positive')
    return result


def distribution(values):
    values = sorted(values)
    return {'count': len(values), 'median': statistics.median(values) if values else None,
            'p95': values[math.ceil(.95 * len(values)) - 1] if values else None,
            'p99': values[math.ceil(.99 * len(values)) - 1] if values else None,
            'percentile_method': 'nearest_rank'}


def collect(binary, destination, mode, scenario, args, metadata):
    name = f'{scenario}-{mode}'
    report_path = destination / f'{name}.json'
    master, slave = pty.openpty()
    events, chunks = [], []
    process = None
    try:
        fcntl.ioctl(slave, termios.TIOCSWINSZ, struct.pack('HHHH', args.rows, args.columns, 0, 0))
        environment = dict(os.environ, TERM='xterm-256color', COLORTERM='truecolor')
        environment.pop('NO_COLOR', None)
        command = [str(binary), '--scenario', scenario, '--mode', mode,
                   '--records', str(args.records), '--points', str(args.points),
                   '--interval', args.interval, '--duration', args.duration,
                   '--warmup', args.warmup, '--report', str(report_path)]
        if args.history:
            command.append('--history')
        if args.headless:
            command.append('--headless')
        process = subprocess.Popen(command, stdin=slave, stdout=slave, stderr=slave,
                                   env=environment, start_new_session=True)
        os.close(slave)
        slave = None
        deadline = time.monotonic() + seconds(args.duration) + seconds(args.warmup) + 30
        eof = False
        while process.poll() is None or not eof:
            if time.monotonic() > deadline:
                raise TimeoutError(f'{name} did not stop after its measurement window')
            if not select.select([master], [], [], .1)[0]:
                continue
            try:
                chunk = os.read(master, 65536)
            except OSError as error:
                if error.errno != errno.EIO:
                    raise
                chunk = b''
            if not chunk:
                eof = True
                if process.poll() is None:
                    process.wait(timeout=5)
                continue
            chunks.append(chunk)
            events.append({'unix_ns': time.time_ns(), 'bytes': len(chunk)})
        if process.wait(timeout=5) != 0:
            raise RuntimeError(f'{name} failed: {b"".join(chunks)[-2000:]!r}')
    finally:
        if process is not None and process.poll() is None:
            process.terminate()
            try:
                process.wait(timeout=5)
            except subprocess.TimeoutExpired:
                process.kill()
                process.wait(timeout=5)
        if slave is not None:
            os.close(slave)
        os.close(master)
        (destination / f'{name}.ansi').write_bytes(b''.join(chunks))
        (destination / f'{name}-output-events.json').write_text(json.dumps(events, indent=2))
    report = json.loads(report_path.read_text())
    report['source_revision'] = metadata['revision']
    report['source_sha256'] = metadata['source_sha256']
    report['build_vcs_matches_source'] = report['revision'] == metadata['revision']
    started, ended = report['measurement_started_unix_ns'], report['measurement_ended_unix_ns']
    if report['measured_elapsed_ns'] < seconds(args.duration) * 1e9 * .9:
        raise RuntimeError(f'{name} ended before its requested measurement window')
    report['terminal_bytes_in_window'] = sum(e['bytes'] for e in events if started <= e['unix_ns'] < ended)
    for index, sample in enumerate(report['samples']):
        following = report['samples'][index + 1]['accepted_unix_ns'] if index + 1 < len(report['samples']) else ended
        first = next((e for e in events if sample['accepted_unix_ns'] <= e['unix_ns'] < following), None)
        sample['first_terminal_output_ms'] = (first['unix_ns'] - sample['accepted_unix_ns']) / 1e6 if first and sample['status'] == 'presented' else None
    report['frame_processing_ms'] = distribution([s['processing_ns'] / 1e6 for s in report['samples'] if s['status'] == 'presented'])
    report['unchanged_processing_ms'] = distribution([s['processing_ns'] / 1e6 for s in report['samples'] if s['status'] == 'unchanged'])
    report['first_output_ms'] = distribution([s['first_terminal_output_ms'] for s in report['samples'] if s['first_terminal_output_ms'] is not None])
    report['cancelled_samples'] = sum(s['status'] == 'cancelled' for s in report['samples'])
    report_path.write_text(json.dumps(report, indent=2))
    return report


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--out', type=Path, required=True)
    parser.add_argument('--scenario', action='append', choices=['unchanged', 'metric', 'request', 'append', 'mixed', 'hidden'])
    parser.add_argument('--mode', action='append', choices=['replace', 'apply'])
    parser.add_argument('--records', type=int, choices=[100, 1000], default=1000)
    parser.add_argument('--points', type=int, choices=[60, 300], default=300)
    parser.add_argument('--interval', default='5s')
    parser.add_argument('--duration', default='2m')
    parser.add_argument('--warmup', default='6s')
    parser.add_argument('--columns', type=int, default=120)
    parser.add_argument('--rows', type=int, default=40)
    parser.add_argument('--history', action='store_true')
    parser.add_argument('--headless', action='store_true')
    args = parser.parse_args()
    for value in [args.interval, args.duration, args.warmup]:
        seconds(value)
    if args.columns < 40 or args.rows < 20:
        parser.error('dashboard requires at least 40 columns and 20 rows')
    if args.headless and (args.columns, args.rows) != (120, 40):
        parser.error('headless mode uses a fixed 120x40 allocation')
    os.chdir(ROOT)
    version = re.search(r'^go (\S+)', (ROOT / 'go.mod').read_text(), re.MULTILINE)[1]
    os.environ.setdefault('GOTOOLCHAIN', 'go' + version)
    destination = args.out.resolve()
    if destination.exists() and any(destination.iterdir()):
        parser.error('out must be a new or empty directory')
    metadata = perf_common.metadata(destination)
    package = json.loads(subprocess.check_output(['go', 'list', '-json', './cmd/perf-periodic'], text=True))
    if Path(package['Dir']).resolve() != ROOT / 'cmd/perf-periodic' or Path(package['Module']['Dir']).resolve() != ROOT:
        raise RuntimeError('Go package source does not match the measured worktree')
    metadata['build_source_directory'] = package['Dir']
    metadata['build_module_directory'] = package['Module']['Dir']
    binary = destination / 'perf-periodic'
    subprocess.run(['go', 'build', '-o', str(binary), './cmd/perf-periodic'], check=True)
    perf_common.verify_sources(metadata)
    metadata['binary_build_info'] = perf_common.output(['go', 'version', '-m', str(binary)])
    metadata['binary_sha256'] = hashlib.sha256(binary.read_bytes()).hexdigest()
    (destination / 'metadata.json').write_text(json.dumps(metadata, indent=2))
    reports = []
    for scenario in args.scenario or ['unchanged', 'mixed']:
        for mode in args.mode or ['replace', 'apply']:
            print(f'Measuring {scenario}/{mode}', flush=True)
            reports.append(collect(binary, destination, mode, scenario, args, metadata))
    perf_common.verify_sources(metadata)
    summary = [{key: report.get(key) for key in ['config', 'headless', 'source_revision', 'source_sha256', 'revision', 'build_vcs_matches_source', 'cpu_available', 'initial_size', 'final_size',
               'measured_elapsed_ns', 'process_cpu_seconds', 'one_core_percent', 'allocated_bytes',
               'allocations', 'gc_count', 'presented_frames_in_window', 'terminal_bytes_in_window',
               'heap_start_after_gc_bytes', 'heap_end_after_gc_bytes', 'frame_processing_ms',
               'unchanged_processing_ms', 'first_output_ms', 'cancelled_samples']} for report in reports]
    (destination / 'summary.json').write_text(json.dumps(summary, indent=2))
    print(destination)


if __name__ == '__main__':
    main()
