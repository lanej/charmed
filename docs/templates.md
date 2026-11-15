## Giving a screen its structure

The application shell assigns space to a header, tabs, content and footer. Supply the content and let the parent coordinate focus and resizing. See the [composition guide](examples.html) for ownership rules and the [molecule catalog](molecules.html) for the individual layout pieces.

For monitoring, start with the [operational dashboard](templates-pattern.html#template/operations). It places summary facts before primary evidence, keeps diagnostics behind inspection, and derives help from active key bindings. Resize and data replacement preserve native viewport customization and independent scroll positions, clamping offsets when content shrinks.

Run `go run ./examples/operations` for the complete application, then use i to inspect, Escape to return, ? to expand help and Ctrl+C to quit. Its [consumer acceptance test](https://github.com/lanej/charmed/blob/main/examples/operations/main_test.go) exercises populated, absent, waiting and failed-refresh observations at wide, standard and narrow allocations.

Compare the saved [overview](assets/operations.png), [narrow](assets/operations-narrow.png), [inspection](assets/operations-details.png), [failed refresh](assets/operations-error.png) and [monochrome](assets/operations-mono.png) frames when adapting the starter. They show deterministic reference observations; capture the consumer's actual output separately before treating its composition as reviewed.
