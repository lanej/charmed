Atomic design

# [Templates](https://lanej.io/charmed/templates.html#level-templates)

Reusable screen structure. Define the layout and ownership of regions before supplying application content.

9 components · 3 groups

[**Layout**  
\
Nested columns2 components](https://lanej.io/charmed/templates-layout.html)

[**Patterns**  
\
Application shell3 components](https://lanej.io/charmed/templates-pattern.html)

[**tui**  
\
Declared screen4 components](https://lanej.io/charmed/templates-tui.html)

## [Giving a screen its structure](https://lanej.io/charmed/templates.html#giving-a-screen-its-structure)

The application shell assigns space to a header, tabs, content and footer. Supply the content and let the parent coordinate focus and resizing. See the [composition guide](https://lanej.io/charmed/examples.html) for ownership rules and the [molecule catalog](https://lanej.io/charmed/molecules.html) for the individual layout pieces.

For monitoring, start with the [operational dashboard](https://lanej.io/charmed/templates-pattern.html#template/operations). It places summary facts before primary evidence, keeps diagnostics behind inspection, and derives help from active key bindings. Resize and data replacement preserve native viewport customization and independent scroll positions, clamping offsets when content shrinks.

Run `go run ./examples/operations` for the complete application, then use i to inspect, Escape to return, ? to expand help and Ctrl+C to quit. Its [consumer acceptance test](https://github.com/lanej/charmed/blob/1a9aa465012af77168ac97c5347d8b3b60184fd4/examples/operations/main_test.go) exercises populated, absent, waiting and failed-refresh observations at wide, standard and narrow allocations.

Compare the saved [overview](https://lanej.io/charmed/assets/operations.png), [narrow](https://lanej.io/charmed/assets/operations-narrow.png), [inspection](https://lanej.io/charmed/assets/operations-details.png), [failed refresh](https://lanej.io/charmed/assets/operations-error.png) and [monochrome](https://lanej.io/charmed/assets/operations-mono.png) frames when adapting the starter. They show deterministic reference observations; capture the consumer's actual output separately before treating its composition as reviewed.
