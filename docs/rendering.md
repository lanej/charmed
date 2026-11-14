# Retained rendering

`runtime.Root` coordinates ordered scene updates and cached presentation. `Root.View` reads the completed frame; cached Git/JSON views require invalidation after direct native presentation edits. Applications can also use Charmed components within their own Bubble Tea models.

## Forward runtime commands

When passing `root` directly to `tea.NewProgram`, Bubble Tea already runs its lifecycle. Applications wrapping a root must forward `Init` and all messages, return every command from `Update`, and let their `View` read the cached frame. Non-input messages include scheduled frames and native command results. Call `Root.Close` to release owned scene resources when the program ends.

```go
type app struct { root *runtime.Root }

func (m *app) Init() tea.Cmd { return m.root.Init() }

func (m *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    _, cmd := m.root.Update(msg)
    return m, cmd
}

func (m *app) View() tea.View { return m.root.View() }
```

State updates are ordered and immediate. Presentation defaults to at most 60 frames per second, with one command pending during bursts and none scheduled while idle; sparse updates do not wait an unnecessary full interval. Bubble Tea owns terminal flushing and its native timer.

## Capture snapshots explicitly

Call `PresentNow` before reading `View` when a synchronous snapshot must include the latest state. `PresentNow` returns errors to the caller and invalidates any pending frame command, so an old timer cannot overwrite the snapshot.

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

Native model access exposes complete viewport content. Git source replacement installs all rows immediately; input bursts can defer row projection until `View`. Full-refresh resize methods reapply presentation; geometry-only methods reuse prepared content. With retained screen or inspector panes, also advance the affected region or pane revision; `Root.Invalidate` does not change application-owned counters.

An optional `PresentationRevision() uint64` on a scene can suppress frames for unchanged updates. Omit it unless the counter advances for every visible change; a data-only counter can hide focus, cursor, scrolling or styling updates.

See [presentation ownership](development.html#batch-native-presentation) for the scheduling contract and [composition](examples.html) for native lifecycle patterns. Keep component updates and presentation on the owning Bubble Tea update thread.
