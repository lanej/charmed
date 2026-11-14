#!/usr/bin/env python3
"""Repeatable native benchmarks. Timings are reports, never shared-runner gates."""
import argparse
import datetime
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import statistics
import subprocess
import tarfile
import tempfile
from perf_common import metadata, verify_sources


def command(args, **kwargs):
    return subprocess.run(args, check=True, text=True, **kwargs)


def read_measurements(source, package):
    measurements = {}
    for line in source.splitlines():
        found = re.match(r'(Benchmark\S+)-\d+\s+\d+\s+([\d.]+) ns/op\s+(\d+) B/op\s+(\d+) allocs/op', line)
        if found:
            key = package + '/' + found[1]
            measurements.setdefault(key, []).append([float(found[2]), int(found[3]), int(found[4])])
    return measurements


def benchmark(args, root, destination, env):
    destination.mkdir(parents=True, exist_ok=True)
    packages = ['story'] if args.suite == 'catalog' else ['layout', 'gitui', 'collection', 'text', 'chart', 'story']
    measurements = {}
    with tempfile.TemporaryDirectory(prefix='charmed-perf-') as temporary:
        for package in packages:
            executable = str(Path(temporary) / (package + '.test'))
            command(['go', 'test', '-c', '-o', executable, './' + package], cwd=root, env=env)
            match = args.match or '^BenchmarkCatalog$'
            if not args.match and args.suite != 'catalog':
                match = '^Benchmark(Render|ResizeRender|PaneResize|Fit|List|Grid|Markdown|JSON|ChartScale|LogIngestion)'
            invocation = [executable, '-test.run', '^$', '-test.bench', match,
                          '-test.benchmem', '-test.benchtime', args.benchtime,
                          '-test.count', str(args.count)]
            if args.suite == 'profiles':
                invocation += ['-test.cpuprofile', str(destination / (package + '.cpu')),
                               '-test.memprofile', str(destination / (package + '.mem'))]
            with (destination / (package + '.txt')).open('w') as output:
                command(invocation, cwd=root, stdout=output, stderr=subprocess.STDOUT, env=env)
            measurements.update(read_measurements((destination / (package + '.txt')).read_text(), package))
    if not measurements:
        raise RuntimeError('No benchmark measurements matched; inspect the raw logs.')
    summary = {key: {'ns_per_op': statistics.median(row[0] for row in rows),
                     'bytes_per_op': statistics.median(row[1] for row in rows),
                     'allocs_per_op': statistics.median(row[2] for row in rows),
                     'samples': rows} for key, rows in measurements.items()}
    (destination / 'results.json').write_text(json.dumps(summary, indent=2) + '\n')
    return summary


def baseline(args, revision, destination, env):
    with tempfile.TemporaryDirectory(prefix='charmed-perf-base-') as temporary:
        root = Path(temporary)
        archive = root / 'source.tar'
        command(['git', 'archive', '--format=tar', '--output', str(archive), revision])
        with tarfile.open(archive) as source:
            source.extractall(root, filter='data')
        overlays = {}
        for package in ['layout', 'gitui', 'collection', 'text', 'chart', 'story']:
            for source in sorted(Path(package).glob('*benchmark_test.go')):
                target = root / source
                target.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(source, target)
                overlays[str(source)] = hashlib.sha256(source.read_bytes()).hexdigest()
        destination.mkdir(parents=True, exist_ok=True)
        (destination / 'baseline.json').write_text(json.dumps({'revision': revision, 'benchmark_overlays': overlays}, indent=2) + '\n')
        return benchmark(args, root, destination, env)


def run():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--suite', choices=['catalog', 'hotspots', 'profiles', 'heap'], default='catalog')
    parser.add_argument('--out', type=Path)
    parser.add_argument('--match', default='')
    parser.add_argument('--count', type=int, default=5)
    parser.add_argument('--benchtime', default='500ms')
    parser.add_argument('--base', help='Compare a committed revision using identical benchmark sources')
    args = parser.parse_args()
    if args.count < 1:
        parser.error('--count must be positive')
    if args.base and args.suite == 'heap':
        parser.error('heap reports are separate observations, not benchmark comparisons')
    if args.out is None:
        stamp = datetime.datetime.now(datetime.timezone.utc).strftime('%Y%m%dT%H%M%S%f')
        args.out = Path('.cache/perf') / (stamp + '-' + args.suite)
    args.out = args.out.resolve()
    if (args.out / 'metadata.json').exists():
        parser.error('Output already contains a report; choose a fresh --out directory.')
    details = metadata(args.out)
    env = dict(os.environ, TERM='xterm-256color', COLORTERM='truecolor', GOWORK='off')
    env.pop('NO_COLOR', None)
    details.update({'suite': args.suite, 'count': args.count, 'benchtime': args.benchtime,
                    'terminal': 'xterm-256color/truecolor', 'match': args.match, 'state': 'started'})
    (args.out / 'metadata.json').write_text(json.dumps(details, indent=2) + '\n')
    with (args.out / 'inventory.json').open('w') as output:
        command(['go', 'run', './cmd/perf-inventory'], stdout=output, env=env)
    if args.suite == 'heap':
        with (args.out / 'heap.json').open('w') as output:
            invocation = ['go', 'run', './cmd/perf-heap']
            if args.match:
                invocation += ['--story', args.match]
            command(invocation, stdout=output, env=env)
        summary = {}
    else:
        before = None
        if args.base:
            revision = subprocess.check_output(['git', 'rev-parse', '--verify', '--end-of-options', args.base + '^{commit}'], text=True).strip()
            before = baseline(args, revision, args.out / 'before', env)
            details['base_revision'] = revision
        summary = benchmark(args, Path.cwd(), args.out / 'after' if before is not None else args.out, env)
        if before is not None:
            comparison = {'added': sorted(summary.keys()-before.keys()), 'removed': sorted(before.keys()-summary.keys()),
                          'workloads': {key: {'before': before[key], 'after': summary[key]}
                                        for key in sorted(before.keys() & summary.keys())}}
            (args.out / 'comparison.json').write_text(json.dumps(comparison, indent=2) + '\n')
    verify_sources(details)
    details['state'] = 'complete'
    (args.out / 'metadata.json').write_text(json.dumps(details, indent=2) + '\n')
    for key, value in sorted(summary.items(), key=lambda item: item[1]['ns_per_op'], reverse=True)[:20]:
        print(f"{value['ns_per_op'] / 1e6:9.3f} ms {value['bytes_per_op']:12,.0f} B {key}")


if __name__ == '__main__':
    run()
