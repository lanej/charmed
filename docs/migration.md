# Migrate to retained presentation

Use this guide when upgrading to the rendering changes in [PR #59](https://github.com/lanej/charmed/pull/59) and [PR #60](https://github.com/lanej/charmed/pull/60). Native `runtime.Scene` signatures remain compatible, but `Root.View` now reads the completed frame, and cached Git/JSON views need explicit invalidation after direct native presentation edits. Applications using their own Bubble Tea root can keep it and apply the component invalidation changes they use.

## Update the dependency

The commit below includes both performance changes and the small-height Git cache repair. Update your application's dependency, retain its previous revision for rollback, and rebuild the binaries that embed Charmed.

```sh
go get github.com/lanej/charmed@e63c5a237b8aac91fa7a6d98d910348697bd328f
go mod tidy
```

## Forward runtime commands

When passing `root` directly to `tea.NewProgram`, Bubble Tea already runs its lifecycle. Applications wrapping a root must forward `Init` and all messages, return every command from `Update`, and let their `View` read the cached frame. Non-input messages include scheduled frames and native command results. Keep existing `Root.Close` cleanup when the program ends.

```go
type app struct { root *runtime.Root }

func (m *app) Init() tea.Cmd { return m.root.Init() }

func (m *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    _, cmd := m.root.Update(msg)
    return m, cmd
}

func (m *app) View() tea.View { return m.root.View() }
```

State updates remain ordered and immediate. Presentation defaults to at most 60 frames per second, with one command pending during bursts and none scheduled while idle; sparse updates do not wait an unnecessary full interval. Bubble Tea still owns terminal flushing and its native timer.

## Capture snapshots explicitly

Replace manual `Update` followed by `View` snapshot sequences with an explicit presentation step. `PresentNow` returns errors to the caller and invalidates any pending frame command, so an old timer cannot overwrite the snapshot.

```go
func snapshot(root *runtime.Root) (tea.View, error) {
    if err := root.PresentNow(); err != nil {
        return tea.View{}, err
    }
    return root.View(), nil
}
```

`PresentNow` does not execute returned application commands. Process those commands and their result messages through the live event loop when their effects are needed before capture. `Render()` is a synchronous convenience; prefer `PresentNow` when snapshot failures must be returned explicitly.

## Invalidate direct native edits

Charmed's own data and interaction methods maintain their presentation dependencies. After directly changing exposed styles or theme fields, invalidate the affected cached component and request a runtime frame from the same serialized application update.

```go
func restyleList(root *runtime.Root, list *gitui.List, style lipgloss.Style) tea.Cmd {
    list.Viewport.Style = style
    list.Invalidate()
    return root.Invalidate()
}
```

| Change | Required action |
| --- | --- |
| Git list native styling or theme | Call `list.Invalidate()`, then return `root.Invalidate()` |
| JSON viewer native styling or theme | Call `viewer.Invalidate()` or `Refresh()`, then return `root.Invalidate()` |
| Direct native viewport scrolling | Return `root.Invalidate()`; viewport offsets already participate in component cache keys |
| Retained screen region contents | Advance the affected `tui.RegionView.Revision`, then request presentation |

Native model access and complete viewport content remain available. Git source replacement still installs all rows immediately; input bursts can defer row projection until `View`. Existing full-refresh and geometry-only resize methods retain their documented distinctions. With retained screen or inspector panes, also advance the affected region or pane revision; `Root.Invalidate` does not change application-owned counters.

An optional `PresentationRevision() uint64` on a scene can suppress frames for unchanged updates. Omit it unless the counter advances for every visible change; a data-only counter can hide focus, cursor, scrolling or styling updates.

Run the application's normal checks, confirm that offline snapshots show the latest state, and exercise small terminal sizes and independent scrolling before rebuilding and deploying. See [presentation ownership](development.html#batch-native-presentation) for the scheduling contract and [composition](examples.html) for existing native lifecycle patterns.
