package main

import (
	"bytes"
	"encoding/json"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

type probeScene struct{ updates int }

func (*probeScene) Init() tea.Cmd            { return nil }
func (s *probeScene) Update(tea.Msg) tea.Cmd { s.updates++; return nil }
func (*probeScene) Resize(layout.Size)       {}
func (*probeScene) View() string             { return "fixture" }

func TestTraceCountsCoalescedInputsOnlyWhenPresented(t *testing.T) {
	var output bytes.Buffer

	root := &runtime.Root{Scene: &probeScene{}, Size: layout.Size{Width: 20, Height: 4}}
	measured := &probe{root: root, trace: json.NewEncoder(&output)}
	measured.Init()
	_, frame := measured.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	measured.Update(tea.KeyPressMsg{Code: tea.KeyUp})

	for range 10 {
		measured.View()
	}

	if output.Len() != 0 {
		t.Fatal("cached reads were reported as completed presentation")
	}

	measured.Update(frame())

	var record frameSample
	if err := json.Unmarshal(output.Bytes(), &record); err != nil {
		t.Fatal(err)
	}

	if record.Kind != "key" || record.Count != 2 || record.Status != "presented" || record.Frame != 2 || record.Duration < record.LatestDuration {
		t.Fatalf("lost coalesced input provenance: %+v", record)
	}

	if measured.pending != [3]pendingSample{} {
		t.Fatal("trace retained completed samples")
	}
}

func TestShutdownReportsAcceptedButUnpresentedInput(t *testing.T) {
	var output bytes.Buffer

	root := &runtime.Root{Scene: &probeScene{}, Size: layout.Size{Width: 20, Height: 4}}
	measured := &probe{root: root, trace: json.NewEncoder(&output)}
	measured.Init()
	measured.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	measured.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})

	var record frameSample
	if err := json.Unmarshal(output.Bytes(), &record); err != nil {
		t.Fatal(err)
	}

	if record.Status != "cancelled" || record.Count != 1 || record.Duration != 0 {
		t.Fatalf("unpresented update lost or mislabeled: %+v", record)
	}
}
