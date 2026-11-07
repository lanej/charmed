import unittest

from preview_pr import BEGIN, END, description, preview_url


class PreviewDescriptionTests(unittest.TestCase):
    def setUp(self):
        self.sha = "a" * 40

    def test_stable_preview_path(self):
        self.assertEqual(
            preview_url(8), "https://lanej.io/charmed/pr-preview/pr-8/"
        )

    def test_preview_leads_before_existing_description(self):
        result = description("## What changed\n\nExisting details.", 8, self.sha)
        self.assertTrue(result.startswith(BEGIN))
        self.assertIn("[Components](https://lanej.io/charmed/pr-preview/pr-8/components.html)", result)
        self.assertIn("[Brand sheet](https://lanej.io/charmed/pr-preview/pr-8/brand.html)", result)
        self.assertIn("## What changed", result)
        self.assertNotIn("github.com/lanej/charmed/pull/8", result)

    def test_update_replaces_existing_block_without_duplicates(self):
        first = description("Details", 8, self.sha)
        updated = description(first, 8, "b" * 40)
        self.assertEqual(updated.count(BEGIN), 1)
        self.assertEqual(updated.count(END), 1)
        self.assertNotIn("aaaaaaaaaaaa", updated)
        self.assertIn("bbbbbbbbbbbb", updated)
        self.assertEqual(description(updated, 8, "b" * 40), updated)

    def test_close_removes_preview_and_retains_content(self):
        initial = "## What changed\n\nDetails."
        self.assertEqual(
            description(description(initial, 8, self.sha), 8, "", remove=True),
            initial,
        )

    def test_invalid_pr_number(self):
        with self.assertRaises(ValueError):
            preview_url(0)


if __name__ == "__main__":
    unittest.main()
