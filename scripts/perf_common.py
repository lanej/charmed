"""Source identity shared by native and microbenchmark reports."""
import datetime
import hashlib
import os
from pathlib import Path
import platform
import subprocess


def output(arguments):
    return subprocess.check_output(arguments, text=True).strip()


def source_state():
    paths = subprocess.check_output(['git', 'ls-files', '--cached', '--others', '--exclude-standard', '-z']).decode().split('\0')
    sources = sorted({p for p in paths if p and (p.endswith(('.go', '.py', '.sh')) or p in ['go.mod', 'go.sum', 'justfile']) and not p.startswith(('.agents/', '.codex/'))})
    manifest = {}
    for name in sources:
        path = Path(name)
        if path.is_file():
            content = path.read_bytes()
            manifest[name] = hashlib.sha256(content).hexdigest()
    digest = hashlib.sha256()
    for name, value in sorted(manifest.items()):
        digest.update((name + '\0' + value + '\n').encode())
    return sources, manifest, digest.hexdigest()


def verify_sources(details):
    if source_state()[2] != details['source_sha256']:
        raise RuntimeError('Measured source changed during the run; report is incomplete. Run again with stable sources.')


def metadata(destination):
    destination = Path(destination)
    root = Path.cwd().resolve()
    destination = destination.resolve()
    try:
        relative = destination.relative_to(root)
    except ValueError:
        relative = None
    if relative is not None and subprocess.run(['git', 'check-ignore', '-q', '--', str(relative)]).returncode != 0:
        raise ValueError('Choose an ignored or out-of-repository performance output directory.')
    destination.mkdir(parents=True, exist_ok=True)
    sources, manifest, identity = source_state()
    patch = subprocess.check_output(['git', 'diff', '--binary', 'HEAD', '--', *sources])
    (destination / 'source.patch').write_bytes(patch)
    untracked = set(subprocess.check_output(['git', 'ls-files', '--others', '--exclude-standard', '-z']).decode().split('\0'))
    for name in sources:
        if name in untracked:
            target = destination / 'new-sources' / name
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_bytes(Path(name).read_bytes())
    cpu = platform.processor() or platform.machine()
    if platform.system() == 'Darwin':
        try:
            cpu = output(['sysctl', '-n', 'machdep.cpu.brand_string'])
        except (OSError, subprocess.CalledProcessError):
            cpu += ' (CPU model unavailable)'
    elif Path('/proc/cpuinfo').exists():
        cpu = next((line.split(':', 1)[1].strip() for line in Path('/proc/cpuinfo').read_text().splitlines()
                    if line.startswith('model name')), cpu)
    return {'revision': output(['git', 'rev-parse', 'HEAD']),
            'dirty': bool(output(['git', 'status', '--porcelain'])),
            'source_sha256': identity, 'source_files': manifest,
            'go': output(['go', 'version']), 'go_env': output(['go', 'env', 'GOTOOLCHAIN', 'GOOS', 'GOARCH']),
            'system': platform.platform(), 'cpu': cpu, 'cpu_count': os.cpu_count(),
            'timestamp': datetime.datetime.now(datetime.timezone.utc).isoformat()}
