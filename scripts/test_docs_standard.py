"""Behavioral examples for the documentation contract and its bypass boundaries."""
from datetime import date
import json
from pathlib import Path
import tempfile
import unittest

from check_docs import Finding, apply_exceptions, check_markdown, check_rendered, words

CONFIG = json.loads((Path(__file__).resolve().parents[1] / 'docs/standards.json').read_text())


class ContentContractTests(unittest.TestCase):
    def check(self, source, profile='tutorial', sections=()):
        return check_markdown(source, 'docs/example.md', CONFIG['profiles'][profile], sections, CONFIG['artifact'])

    def rules(self, source):
        return {finding.rule for finding in self.check(source)}

    def test_visible_words_include_code_but_exclude_link_destinations(self):
        self.assertEqual(words('Read [the guide](https://example.com/a/very/long/path) with `SetData`.'), 5)

    def test_headings_do_not_reset_the_prose_budget(self):
        paragraph = ' '.join(['observation'] * 70)
        findings = self.check(('\n\n## Another heading\n\n'.join([paragraph] * 4)))
        self.assertIn('run.words', {finding.rule for finding in findings})

    def test_splitting_into_short_paragraphs_fails_lower_and_mean_bounds(self):
        source = '\n\n'.join(['Small fragments make reading harder.'] * 6)
        rules = self.rules(source)
        self.assertTrue({'paragraph.words', 'run.blocks', 'run.mean_words'} <= rules)

    def test_lists_contribute_to_uninterrupted_word_budget(self):
        source = '\n'.join(['- ' + ' '.join(['item'] * 70)] * 4)
        self.assertIn('run.words', self.rules(source))

    def test_link_only_rows_still_count_toward_run_budgets(self):
        row = '[A guide with descriptive labels](guide.html)'
        self.assertIn('run.blocks', self.rules('\n\n'.join([row] * 6)))

    def test_oversized_paragraph_has_location_and_correction(self):
        findings = self.check('# Title\n\n' + ' '.join(['word'] * 101))
        finding = next(f for f in findings if f.rule == 'paragraph.words')
        self.assertEqual((finding.line, finding.actual), (3, 101))
        self.assertIn('18–90', str(finding))
        self.assertIn('Combine', str(finding))

    def test_meaningful_code_resets_a_text_run(self):
        paragraph = ' '.join(['measurement'] * 80)
        self.assertEqual(self.check(paragraph + '\n\n```go\nplot.View()\n```\n\n' + paragraph), [])

    def test_adjacent_preview_and_code_need_no_padding_prose(self):
        self.assertEqual(self.check('```text\nRequests 72\n```\n\n```go\nplot.View()\n```'), [])

    def test_empty_or_comment_only_fences_are_not_artifacts(self):
        self.assertIn('artifact.code', self.rules('```go\n// a decoration\n```'))
        self.assertIn('artifact.code', self.rules('```go\n```'))

    def test_decorative_image_does_not_reset_budget(self):
        paragraph = ' '.join(['measurement'] * 85)
        source = paragraph + '\n\n![icon](image.png)\n\n' + paragraph + '\n\n' + paragraph
        self.assertTrue({'artifact.image', 'run.words'} <= self.rules(source))

    def test_meaningful_table_resets_budget(self):
        paragraph = ' '.join(['measurement'] * 80)
        table = '| State | Meaning |\n| --- | --- |\n| Empty | No observations |\n| Known | Supplied observations |'
        self.assertEqual(self.check(paragraph + '\n\n' + table + '\n\n' + paragraph), [])

    def test_one_row_table_does_not_reset_budget(self):
        paragraph = ' '.join(['measurement'] * 85)
        table = '| State | Meaning |\n| --- | --- |\n| Empty | No observations |'
        self.assertIn('run.words', self.rules(paragraph + '\n\n' + table + '\n\n' + paragraph + '\n\n' + paragraph))

    def test_table_cannot_hide_an_oversized_paragraph(self):
        table = '| State | Meaning |\n| --- | --- |\n| Empty | ' + ' '.join(['detail'] * 71) + ' |\n| Known | Observed |'
        self.assertIn('table.cell_words', self.rules(table))

    def test_required_sections_are_profile_specific(self):
        self.assertIn('section.required', {f.rule for f in self.check('# A page', sections=['Install'])})
        self.assertEqual(self.check('# A page\n\n## Install', sections=['Install']), [])

    def test_code_size_has_an_upper_bound(self):
        self.assertIn('code.lines', self.rules('```go\n' + 'plot.View()\n' * 49 + '```'))

    def test_exact_reviewed_exception_and_expiration(self):
        finding = Finding('page.md', 3, 'paragraph.words', 101, '18–90', 'Rewrite')
        exception = dict(path='page.md', line=3, rule='paragraph.words', actual=101, reason='Upstream protocol text', owner='maintainer', review_by='2030-02-02')
        self.assertEqual(apply_exceptions([finding], [exception], date(2030, 2, 1)), [])
        with self.assertRaises(ValueError):
            apply_exceptions([finding], [exception], date(2030, 2, 2))
        with self.assertRaises(ValueError):
            apply_exceptions([], [exception], date(2030, 2, 1))

    def test_generated_links_and_empty_output_are_rejected(self):
        with tempfile.TemporaryDirectory() as directory:
            site = Path(directory)
            (site / 'index.html').write_text('<main><a href="missing.html#bad">Link</a><section class="component"><div role="tabpanel" id="panel-example"><pre class="terminal"></pre></div></section></main>')
            findings = check_rendered(site, CONFIG)
            rules = {f.rule for f in findings}
            self.assertTrue({'link.local', 'component.section', 'preview.empty'} <= rules)

    def test_heatmap_cells_are_visible_even_without_text_glyphs(self):
        with tempfile.TemporaryDirectory() as directory:
            site = Path(directory)
            (site / 'index.html').write_text('<section class="component"><div role="tabpanel" id="panel-example"><pre class="terminal"><span style="background-color:#4e70a1">    </span></pre></div></section>')
            self.assertNotIn('preview.empty', {f.rule for f in check_rendered(site, CONFIG)})

    def test_duplicate_anchor_and_missing_target_anchor_are_rejected(self):
        with tempfile.TemporaryDirectory() as directory:
            site = Path(directory)
            (site / 'index.html').write_text('<main id="same"><p id="same">Duplicate</p><a href="#absent">Invalid anchor</a></main>')
            self.assertTrue({'link.local', 'link.duplicate_id'} <= {f.rule for f in check_rendered(site, CONFIG)})


if __name__ == '__main__':
    unittest.main()
