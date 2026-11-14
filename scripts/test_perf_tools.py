"""Contract tests for measurement summaries and native resource accounting."""
import unittest
from unittest.mock import patch
from perf_native import distribution, process_stats, linux_stats
from perf import read_measurements


class PerformanceReportTests(unittest.TestCase):
    def test_empty_samples_remain_missing(self):
        self.assertEqual(distribution([]), {'samples': 0})

    def test_quantiles_and_count_preserve_samples(self):
        result = distribution(list(range(1, 101)))
        self.assertEqual(result['samples'], 100)
        self.assertEqual(result['median_ns'], 50.5)
        self.assertEqual(result['p95_ns'], 95)
        self.assertEqual(result['p99_ns'], 99)

    @patch('perf_native.platform.system', return_value='Darwin')
    @patch('perf_native.subprocess.check_output', return_value='  1:02.34  2048\n')
    def test_process_stats_units(self, command, system):
        self.assertEqual(process_stats(123), (62.34, 2097152))
        command.assert_called_once_with(['ps', '-p', '123', '-o', 'time=', '-o', 'rss='], text=True)

    @patch('perf_native.platform.system', return_value='Darwin')
    @patch('perf_native.subprocess.check_output', return_value='  01:02:03.45  1024\n')
    def test_long_process_cpu_time(self, command, system):
        self.assertEqual(process_stats(123), (3723.45, 1048576))

    @patch('perf_native.os.sysconf', side_effect=lambda name: 100 if name == 'SC_CLK_TCK' else 4096)
    def test_linux_ticks_and_resident_pages(self, configuration):
        fields = ['0'] * 22
        fields[0], fields[11], fields[12], fields[21] = 'S', '123', '77', '512'
        self.assertEqual(linux_stats('123 (name with spaces) ' + ' '.join(fields)), (2.0, 2097152))

    def test_benchmark_parser_keeps_repetitions_and_units(self):
        raw = "BenchmarkSmall/label-12 100 12.5 ns/op 64 B/op 2 allocs/op\nBenchmarkSmall/label-12 100 13 ns/op 72 B/op 3 allocs/op\nPASS\n"
        self.assertEqual(read_measurements(raw, 'text'), {'text/BenchmarkSmall/label': [[12.5, 64, 2], [13.0, 72, 3]]})

    @patch('perf_common.subprocess.run')
    def test_report_destination_must_not_pollute_sources(self, command):
        from pathlib import Path
        from perf_common import metadata
        command.return_value.returncode = 1
        with self.assertRaisesRegex(ValueError, 'ignored or out-of-repository'):
            metadata(Path.cwd() / 'unignored-performance-output')

    @patch('perf_common.source_state', return_value=([], {}, 'changed'))
    def test_source_drift_is_not_a_successful_report(self, state):
        from perf_common import verify_sources
        with self.assertRaisesRegex(RuntimeError, 'source changed'):
            verify_sources({'source_sha256': 'original'})
