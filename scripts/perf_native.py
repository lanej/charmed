#!/usr/bin/env python3
"""Native PTY measurements: view latency, first output, idle CPU and resident memory."""
import argparse
import datetime
import errno
import fcntl
import json
import math
import os
import platform
from pathlib import Path
import pty
import select
import signal
import statistics
import struct
import subprocess
import tempfile
import termios
import threading
import time
from perf_common import metadata, verify_sources


def linux_stats(source):
    fields = source.rpartition(') ')[2].split()
    ticks = os.sysconf('SC_CLK_TCK')
    return (int(fields[11])+int(fields[12]))/ticks, int(fields[21])*os.sysconf('SC_PAGE_SIZE')


def process_stats(pid):
    if platform.system() == 'Linux':
        return linux_stats(Path(f'/proc/{pid}/stat').read_text())
    raw = subprocess.check_output(['ps', '-p', str(pid), '-o', 'time=', '-o', 'rss='], text=True).split()
    total = 0.0
    for part in raw[0].split(':'):
        total = total * 60 + float(part)
    return total, int(raw[1]) * 1024


def cpu_resolution():
    if platform.system() == 'Linux':
        return 1/os.sysconf('SC_CLK_TCK')
    value = subprocess.check_output(['ps', '-p', str(os.getpid()), '-o', 'time='], text=True).strip()
    return 10**(-len(value.rsplit('.', 1)[1])) if '.' in value else 1


def distribution(values):
    if not values:
        return {'samples': 0}
    ordered = sorted(values)
    return {'samples': len(values), 'median_ns': statistics.median(values),
            'p95_ns': ordered[min(len(ordered)-1, max(0, math.ceil(len(ordered)*.95)-1))],
            'p99_ns': ordered[min(len(ordered)-1, max(0, math.ceil(len(ordered)*.99)-1))]}


def measure(binary, fixture, trace, count, idle, live=False):
    master, slave = pty.openpty()
    def resize(width, height):
        fcntl.ioctl(master, termios.TIOCSWINSZ, struct.pack('HHHH', height, width, 0, 0))
    resize(120, 40)
    env = dict(os.environ, TERM='xterm-256color', COLORTERM='truecolor')
    env.pop('NO_COLOR', None)
    invocation = [binary, '--story', fixture, '--trace', str(trace)]
    if live:
        invocation += ['--live']
    try:
        process = subprocess.Popen(invocation, stdin=slave, stdout=slave, stderr=slave, env=env, start_new_session=True)
    except BaseException:
        os.close(master)
        raise
    finally:
        os.close(slave)
    stopped = threading.Event()
    condition = threading.Condition()
    output = {'bytes': 0, 'last': 0, 'first': 0, 'waiting': False, 'error': None}
    def drain():
        while not stopped.is_set():
            try:
                if not select.select([master], [], [], .05)[0]:
                    continue
                data = os.read(master, 65536)
            except OSError as error:
                if error.errno == errno.EIO:
                    return
                with condition:
                    output['error'] = str(error)
                    condition.notify_all()
                return
            if not data:
                return
            with condition:
                output['bytes'] += len(data)
                output['last'] = time.monotonic_ns()
                if output['waiting'] and not output['first']:
                    output['first'] = output['last']
                condition.notify_all()
    reader = threading.Thread(target=drain)
    try:
        reader.start()
        time.sleep(1.5)
        warm_events = len(trace.read_text().splitlines())
        cpu, rss = process_stats(process.pid)
        before, started = output['bytes'], time.monotonic()
        time.sleep(idle)
        elapsed = time.monotonic() - started
        after_cpu, after_rss = process_stats(process.pid)
        result = {'story': fixture, 'live': live, 'idle_seconds': elapsed,
                  'idle_cpu_seconds': after_cpu-cpu, 'idle_core_percent': 100*(after_cpu-cpu)/elapsed,
                  'idle_terminal_bytes': output['bytes']-before, 'rss_before_bytes': rss,
                  'rss_after_idle_bytes': after_rss}
        latencies, timeouts, memory = [], 0, []
        def sample(action):
            with condition:
                output['first'], output['waiting'] = 0, True
                sent = time.monotonic_ns()
                action()
                deadline = time.monotonic() + 2
                while not output['first'] and time.monotonic() < deadline:
                    condition.wait(timeout=.05)
                latency = output['first']-sent if output['first'] else None
                output['waiting'] = False
            time.sleep(.06)
            return latency
        def resize_child(width, height):
            resize(width, height)
            # This owned child has a new session but no controlling foreground
            # terminal. Deliver SIGWINCH explicitly after changing PTY geometry.
            os.kill(process.pid, signal.SIGWINCH)
        if not live and fixture != 'screen/services':
            for index in range(count):
                latency = sample(lambda: resize_child(120 if index % 2 else 80,
                                                     40 if index % 2 else 24))
                if latency is None:
                    timeouts += 1
                else:
                    latencies.append(latency)
                if index % 10 == 0:
                    memory.append(process_stats(process.pid)[1])
        key_latencies, key_timeouts = [], 0
        if fixture == 'git/browser':
            for index in range(count):
                latency = sample(lambda: os.write(master, b'\x1b[A' if index % 2 else b'\x1b[B'))
                if latency is None:
                    key_timeouts += 1
                else:
                    key_latencies.append(latency)
        result['key_first_output'] = distribution(key_latencies)
        result['key_no_output_samples'] = key_timeouts
        result['resize_first_output'] = distribution(latencies)
        result['resize_no_output_samples'] = timeouts
        result['rss_resize_samples_bytes'] = memory
    finally:
        try:
            if process.poll() is None:
                try:
                    os.write(master, b'\x03')
                except OSError:
                    process.terminate()
                try:
                    process.wait(timeout=3)
                except subprocess.TimeoutExpired:
                    process.kill()
                    process.wait(timeout=3)
        finally:
            stopped.set()
            if reader.ident is not None:
                reader.join(timeout=3)
            os.close(master)
    if output['error']:
        raise RuntimeError(output['error'])
    if process.returncode != 0:
        raise RuntimeError(f'{fixture}: native probe exited {process.returncode}')
    events = [json.loads(line) for line in trace.read_text().splitlines()][warm_events:]
    if not live and fixture != 'screen/services':
        resize_events = sum(e['kind'] == 'resize' for e in events)
        if resize_events < count:
            raise RuntimeError(f'{fixture}: expected {count} resize events, observed {resize_events}')
    if fixture == 'git/browser' and sum(e['kind'] == 'key' for e in events) < count:
        raise RuntimeError('Git probe did not process the expected key sequence')
    result['update_to_view'] = {kind: distribution([e['duration_ns'] for e in events if e['kind'] == kind])
                                for kind in ['key', 'resize', 'message']}
    return result


def run():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--out', type=Path)
    parser.add_argument('--samples', type=int, default=100)
    parser.add_argument('--idle-seconds', type=float, default=5)
    args = parser.parse_args()
    if args.samples < 1 or args.idle_seconds <= 0:
        parser.error('samples and idle seconds must be positive')
    if args.out is None:
        stamp = datetime.datetime.now(datetime.timezone.utc).strftime('%Y%m%dT%H%M%S%f')
        args.out = Path('.cache/perf') / (stamp + '-native')
    if (args.out / 'metadata.json').exists():
        parser.error('Output already contains a report; choose a fresh --out directory.')
    details = metadata(args.out)
    details['state'] = 'started'
    (args.out / 'metadata.json').write_text(json.dumps(details, indent=2) + '\n')
    results = []
    with tempfile.TemporaryDirectory(prefix='charmed-native-') as temporary:
        binary = str(Path(temporary) / 'probe')
        subprocess.run(['go', 'build', '-o', binary, './cmd/perf-probe'], check=True)
        for fixture, live in [('git/browser', False), ('text/markdown', False),
                              ('screen/services', False), ('feedback/spinner', True), ('feedback/time', True)]:
            trace = args.out / (fixture.replace('/', '-') + '.jsonl')
            results.append(measure(binary, fixture, trace, args.samples, args.idle_seconds, live))
    verify_sources(details)
    details['state'] = 'complete'
    (args.out / 'metadata.json').write_text(json.dumps(details, indent=2) + '\n')
    report = {'metadata': details, 'cpu_counter_resolution_seconds': cpu_resolution(),
              'limits': 'Instrumented update-to-view and first PTY bytes, not completed terminal painting. RSS is not retained Go heap. Initial warmup events are excluded.',
              'results': results}
    (args.out / 'native.json').write_text(json.dumps(report, indent=2) + '\n')
    print(json.dumps(report, indent=2))


if __name__ == '__main__':
    run()
