// Command perf-heap samples retained Go heap across repeated native resizes.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	goruntime "runtime"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/story"
)

type observation struct {
	ID         string   `json:"id"`
	Heap       []uint64 `json:"retained_heap_bytes"`
	AfterClose uint64   `json:"after_close_heap_bytes"`
}

func heap() uint64 {
	goruntime.GC()

	var stats goruntime.MemStats
	goruntime.ReadMemStats(&stats)

	return stats.HeapAlloc
}

func measure(id string, cycles, phases int) (observation, error) {
	definition, err := story.Find(id)
	if err != nil {
		return observation{}, err
	}

	model, err := story.New(id, definition.Defaults, nil)
	if err != nil {
		return observation{}, err
	}

	exercise := func() {
		for _, size := range [][2]int{{40, 12}, {120, 40}, {200, 60}} {
			_, _ = model.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
			model.Render()
		}
	}
	exercise()

	result := observation{ID: id, Heap: []uint64{heap()}}

	for range phases {
		for range cycles {
			exercise()
		}

		result.Heap = append(result.Heap, heap())
	}

	err = model.Close()
	model = nil
	result.AfterClose = heap()

	return result, err
}

func run() error {
	id := flag.String("story", "all", "Fixture ID or all")
	cycles := flag.Int("cycles", 20, "Resize cycles per phase")
	phases := flag.Int("phases", 4, "Post-warmup phases")

	flag.Parse()

	if *cycles < 1 || *phases < 1 {
		return errors.New("cycles and phases must be positive")
	}

	ids := []string{*id}
	if *id == "all" {
		ids = nil
		for _, definition := range story.Definitions() {
			ids = append(ids, definition.ID)
		}
	}

	result := []observation{}

	for _, name := range ids {
		observed, err := measure(name, *cycles, *phases)
		if err != nil {
			return err
		}

		result = append(result, observed)
	}

	return json.NewEncoder(os.Stdout).Encode(result)
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
