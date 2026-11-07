# Compose behavior as well as appearance

A complete screen owns one Go model and coordinates its child widgets. The service workbench links an input, table, time-series chart and details viewport. The deployment example coordinates host distributions, histories and inspection; [read its telemetry contract and interpretation](deployment.html).

## Ownership rules

| Owner | Responsibility |
| --- | --- |
| Theme | Semantic colors, contrast, ASCII and motion policy |
| Layout parent | Allocate rectangles in columns and rows, including borders |
| Focus scope | Route keyboard input to the active child |
| Component | Own its native widget state and return commands |
| Root runtime | Own the terminal session and forward commands |
| Application | Supply data and handle typed UI events |

In the native service screen, Tab changes focus, arrows select rows and Ctrl+N advances background samples. Selection updates the chart and inspector together. Resize preserves selection; hidden regions leave the focus order.

In the native deployment screen, arrows choose hosts and windows, brackets choose a response class, 1/2/3 change displayed spread, and b switches between stacked and disconnected bands. These are terminal controls; the saved frames below do not execute them.

The deployment fixture’s saved variants include the disconnected band presentation, rendered on the same night surface as the default snapshot.

## Screen fixtures
