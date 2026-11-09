package diagram

import "fmt"

// TransitionTrace records transition indices in chronological order, oldest first.
// Indices identify exact edges, including parallel edges. Last defaults to one;
// All shows the whole trace instead. The renderer never executes transitions.
type TransitionTrace struct {
	Steps []int
	Last  int
	All   bool
}

func (d StateDiagram) traceSteps() ([]int, error) {
	t := d.Trace
	if t.Last < 0 || t.All && t.Last > 0 {
		return nil, fmt.Errorf("trace Last must be nonnegative and cannot be combined with All")
	}

	for i, step := range t.Steps {
		if step < 0 || step >= len(d.Transitions) {
			return nil, fmt.Errorf("trace step %d references unknown transition %d", i, step)
		}

		if i > 0 && d.Transitions[t.Steps[i-1]].To != d.Transitions[step].From {
			return nil, fmt.Errorf("trace step %d does not continue the previous transition", i)
		}
	}

	if len(t.Steps) == 0 {
		return nil, nil
	}

	if d.Active != "" && d.Transitions[t.Steps[len(t.Steps)-1]].To != d.Active {
		return nil, fmt.Errorf("trace does not end at active state %q", d.Active)
	}

	last := max(1, t.Last)
	if t.All {
		last = len(t.Steps)
	}

	return t.Steps[max(0, len(t.Steps)-last):], nil
}
