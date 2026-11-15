# Operational dashboard transfer experiment

A fresh-context agent built the supplied backup-verification consumer using Charmed's operational dashboard, bucket-count component and consumer workflow. A separate reviewer examined raw source, native frames and interaction evidence before reading the builder's handoff. The first result was frozen before review; no review-driven consumer repair was made.

The bounded trial used one synthetic task, one model (gpt-6.1-sol, medium effort), and one 30-minute implementation allowance. It passed the scoped semantic, keyboard, responsive-layout and measured state-retention checks. This is one successful transfer example, not an A/B comparison, measurement of time saved, general effectiveness result or user aesthetic approval.

## Run the retained consumer

The consumer source and fixture are unchanged from that first result; [source-fingerprint.json](source-fingerprint.json) records their hashes. They now compile as part of Charmed's root module and its normal quality gate, without a separate dependency installation or machine-specific environment script.

```sh
go run ./experiments/charmed-tui-transfer/consumer \
  -fixture experiments/charmed-tui-transfer/consumer/fixture.json
gotestsum -- -race -count=1 ./experiments/charmed-tui-transfer/consumer
```

Use 1–5 to select waiting, confirmed empty, ready, failed refresh or replacement. Use i/Escape to inspect/return, j/k to scroll, r to reapply the current observation, ? for help and q to quit. `-ascii` and `-mono` are separate presentation preferences. The fixture's UTC clock is fixed; no backup service, scheduler or network client is invoked.

## Task and observations

The operator must distinguish campaign completion from verification health and interval history. Ready is 47/60, replacement is 48/60, while numeric interval history totals 45 and complete-only history totals 42. Complete interval counts include 40, 0 and 2; partial is 3, and missing/unknown intervals have no count. Region availability, health and last verification time are independent. A failed refresh retains the prior observation instead of presenting an attempted refresh as success.

| Check | Observed result |
| --- | --- |
| Exact evidence and refresh states | Campaign counts, zero/small values, interval coverage and region distinctions remain explicit; waiting differs from confirmed empty |
| Native interaction and responsive layouts | Actual app keys, resize and normal quit worked; content remained reachable at 132x40, 80x24 and 40x18 |
| Independent scroll positions | Nonzero overview/inspection offsets 4/8 survived refresh and back/reopen while still valid |
| Replacement clamping | At 40x18, inspection shrank from 70 lines at offset 60 to 59 lines with maximum 49, and clamped to 49 |
| Independent replay | All 15 original ANSI frames reproduced exactly from frozen source; native interaction evidence was checked separately |
| Rendered review | 17 Unicode, ASCII and monochrome images were inspected; the wide view retains lower blank space and narrow views require scrolling |

Saved consumer output: [wide overview](images/overview-132x40.png), [narrow overview](images/overview-40x18.png), [failed refresh](images/failed-refresh-80x24.png), and [failed refresh in ASCII](images/failed-ascii.png). These are rendered native frames, not recreated HTML interfaces; browser projections supplement native cell and interaction checks rather than establishing every physical font's behavior.

## Scope and follow-up

The trial used a working dependency snapshot based on `d7f55a5b1ca927c810b085327fdb9bbf154b791a`. The workflow was explicitly loaded as a standalone skill; it is now integrated into the existing public Charmed skill package. Current-main integration and package checks are separate validation, not a second agent trial or a study of automatic skill discovery. Raw machine-specific logs, private tool transcripts and dependency archives are retained locally rather than published here.

Supplemental review found an inherited library defect: ASCII mode removes semantic warning color while accent/chart colors remain. The failure label stays explicit, so the task's semantic fallback check passes, but glyph/color independence is incomplete. [Issue #66](https://github.com/lanej/charmed/issues/66) records that defect; it was not silently fixed in the experiment or this feature.
