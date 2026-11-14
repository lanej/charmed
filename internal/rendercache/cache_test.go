package rendercache

import (
	"strings"
	"testing"
	"unsafe"
)

func TestCacheTracksDependenciesAndOwnsCurrentOutput(t *testing.T) {
	type key struct {
		revision uint64
		width    int
	}

	var cache Cache[key]

	calls := 0
	source := strings.Repeat("large source", 10000)
	render := func() string { calls++; return source[:3] }

	first := cache.Get(key{1, 80}, render)
	//nolint:gosec // G103: compare immutable backing addresses without dereferencing or mutating them.
	if unsafe.StringData(first) == unsafe.StringData(source) {
		t.Fatal("short output retained source backing storage")
	}

	for range 20 {
		if cache.Get(key{1, 80}, render) != first {
			t.Fatal("warm output changed")
		}
	}

	if calls != 1 {
		t.Fatal("same dependencies recomputed output")
	}

	cache.Get(key{1, 40}, render)
	cache.Get(key{2, 40}, render)

	if calls != 3 {
		t.Fatal("changed dependency reused stale output")
	}

	if allocations := testing.AllocsPerRun(10, func() { cache.Get(key{2, 40}, render) }); allocations != 0 {
		t.Fatalf("cached read allocated %v times", allocations)
	}

	cache.Clear()

	if cache.value != "" || cache.valid {
		t.Fatal("clear retained output")
	}
}
