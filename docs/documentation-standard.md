# Documentation as a designed interface

Every page should help someone decide what to use and produce a useful result. The versioned [standard configuration](https://github.com/lanej/charmed/blob/main/docs/standards.json) defines bounded variation for each page type; the shared component template supplies a consistent reading order.

## Content rhythm

Let W be the visible words in a text run and P its prose blocks. A run crosses headings and blank lines until a meaningful artifact. Both W and P have bounds; W/P has a minimum when P exceeds one, discouraging piles of small paragraphs.

| Profile | Words per paragraph | Words between artifacts | Maximum prose blocks | Minimum W/P |
| --- | --- | --- | --- | --- |
| Overview | 12–90 | 12–240 | 4 | 20 |
| Tutorial | 18–90 | 18–240 | 4 | 24 |
| Recipe | 14–100 | 14–300 | 4 | 24 |
| Reference | 14–100 | 14–360 | 5 | 24 |
| Contributor | 14–100 | 14–300 | 4 | 24 |

Words include inline code and link labels, excluding destinations and formatting. List content contributes to the same run budget. Headings never reset it. Adjacent artifacts may need no intervening prose: a native preview followed by its usage code is an intentional pair.

A code block must contain useful code, an image needs at least six descriptive alternative-text words, and a comparison table needs two data rows and two columns. Captions, decorative images and empty blocks cannot reset the budget. Combine related explanation before adding another section or artifact.

| Component contract | Accepted range or requirement |
| --- | --- |
| When to use it | One paragraph of 6–65 words |
| Minimal consumer example | 1–48 displayed lines, with explicit data and no fixture dependencies |
| Updates, focus and sizing | One paragraph of 28–100 words |
| Example output | 1–36 native rows, with meaningful initialized content |
| Runnable source | A linked Go file plus the exact catalog-render command |

## Rendered layout

Templates place the purpose, native example, usage code, lifecycle guidance and run command in that order. Prose stays within a readable column; wide native output retains terminal geometry. Code remains fully visible vertically, with horizontal scrolling confined to its block when necessary.

| Rendered measurement | Accepted range |
| --- | --- |
| Prose measure on a wide allocation | 28–72 character widths |
| Component heading to first useful preview | 0–540 pixels |
| Default example height | 20–600 pixels |
| Gap between preview and usage | 0–180 pixels |
| Section spacing | 20–64 pixels |

Chromium measures these bounds at mobile and desktop widths, alongside overflow, selected examples and navigation. Source compilation cannot establish visual usefulness: reviewers must inspect whether the chosen data and framing demonstrate the component. The file picker must show loaded files, and chart axes must make the plotted observations understandable.

## Review and exceptions

Run the fast source gate while editing and install the staged-tree commit hook once per checkout. CI runs the same source gate, validates generated local links and component structure, then measures the browser layout. Skipping a local hook does not bypass those checks.

```sh
just docs-check
just hooks
just docs
```

Failures report a file, line, rule, observed value, accepted range and suggested correction. Bounds live in one configuration. A reviewed exception must identify an exact finding, reason, owner and future review date; stale, ambiguous or expired exceptions fail. There are no initial exceptions or blanket suppressions.

Review the [chart](atoms-chart.html#chart/timeseries), [input](atoms-input.html#input/text) and [inspector](organisms-logs.html#pattern/log-inspector) together when changing the template. They cover static data, editable state and a coordinated composition. Retain representative mobile and desktop screenshots in the PR, then review the published preview before merging.
