import argparse
import importlib.util
from pathlib import Path
import unittest

spec = importlib.util.spec_from_file_location('periodic_measure', Path(__file__).with_name('measure.py'))
measure = importlib.util.module_from_spec(spec)
spec.loader.exec_module(measure)


class MeasurementTests(unittest.TestCase):
    def test_durations(self):
        self.assertEqual(measure.seconds('5s'), 5)
        self.assertEqual(measure.seconds('2m'), 120)
        self.assertEqual(measure.seconds('100ms'), .1)
        for invalid in ['0s', '-1s', 'never', '5']:
            with self.assertRaises(argparse.ArgumentTypeError):
                measure.seconds(invalid)

    def test_absent_output_has_no_zero_latency(self):
        result = measure.distribution([])
        self.assertEqual(result['count'], 0)
        self.assertIsNone(result['median'])
        self.assertIsNone(result['p95'])

    def test_sample_counts_and_nearest_rank(self):
        result = measure.distribution([3, 1, 2, 4])
        self.assertEqual(result['count'], 4)
        self.assertEqual(result['median'], 2.5)
        self.assertEqual(result['p95'], 4)
        self.assertEqual(result['p99'], 4)


if __name__ == '__main__':
    unittest.main()
