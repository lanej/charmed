// Package story is the single source of truth for native and MDX component fixtures.
package story

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/charmbracelet/colorprofile"
	"math"
	"slices"

	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/theme"
)

type Config struct {
	Columns           int        `json:"columns"`
	Rows              int        `json:"rows"`
	Theme             theme.Mode `json:"theme"`
	Profile           string     `json:"profile"`
	ASCII             bool       `json:"ascii"`
	ReducedMotion     bool       `json:"reducedMotion"`
	Live              bool       `json:"live"`
	State             string     `json:"state"`
	Value             string     `json:"value"`
	Series            int        `json:"series"`
	Amount            float64    `json:"amount"`
	Horizontal        bool       `json:"horizontal"`
	Sigma             int        `json:"sigma"`
	Scenario          string     `json:"scenario"`
	ReferenceID       string     `json:"referenceID"`
	CandidateID       string     `json:"candidateID"`
	Boundaries        bool       `json:"boundaries"`
	ComparisonVariant string     `json:"comparisonVariant"`
	Emphasis          string     `json:"emphasis"`
}

func DefaultConfig() Config {
	return Config{Columns: 100, Rows: 30, Theme: theme.Dark, Profile: "truecolor", State: "ready", Value: "production", Series: 2, Amount: 68, Sigma: 3, Scenario: "isolated-hosts", ReferenceID: fixture.ReferenceVersionID, CandidateID: fixture.CandidateVersionID, ComparisonVariant: "stacked", Emphasis: "normal"}
}
func (c Config) Validate() error {
	if !slices.Contains([]string{"", "stacked", "bands"}, c.ComparisonVariant) {
		return errors.New("comparison variant must be stacked or bands")
	}
	if !slices.Contains([]string{"", "bright", "normal", "dim"}, c.Emphasis) {
		return errors.New("emphasis must be bright, normal or dim")
	}
	if c.Columns < 20 || c.Columns > 240 || c.Rows < 8 || c.Rows > 80 {
		return errors.New("terminal size must be 20–240 columns and 8–80 rows")
	}
	if !slices.Contains([]theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono}, c.Theme) {
		return fmt.Errorf("unknown theme %q", c.Theme)
	}
	if !slices.Contains([]string{"truecolor", "ansi256", "ansi16", "none"}, c.Profile) {
		return fmt.Errorf("unknown color profile %q", c.Profile)
	}
	if !slices.Contains([]string{"ready", "empty", "loading", "error", "disabled"}, c.State) {
		return fmt.Errorf("unknown fixture state %q", c.State)
	}
	if len(c.Value) > 4096 {
		return errors.New("fixture value is limited to 4096 bytes")
	}
	if c.Series < 1 || c.Series > 4 {
		return errors.New("series must be between 1 and 4")
	}
	if c.Sigma < 1 || c.Sigma > 3 {
		return errors.New("sigma must be 1, 2 or 3")
	}
	if err := (fixture.Versions{Reference: c.ReferenceID, Candidate: c.CandidateID}).Validate(); err != nil {
		return err
	}
	if !slices.Contains([]string{"isolated-hosts", "uniform-regression", "unequal-traffic", "insufficient-samples"}, c.Scenario) {
		return fmt.Errorf("unknown scenario %q", c.Scenario)
	}
	if math.IsNaN(c.Amount) || math.IsInf(c.Amount, 0) || c.Amount < 0 || c.Amount > 100 {
		return errors.New("amount must be between 0 and 100")
	}
	return nil
}
func (c Config) Size() layout.Size { return layout.Size{Width: c.Columns, Height: c.Rows} }
func (c Config) Palette() theme.Theme {
	t := theme.New(c.Theme)
	t.ASCII = t.ASCII || c.ASCII
	t.ReducedMotion = t.ReducedMotion || c.ReducedMotion
	return t
}
func (c Config) ColorProfile() colorprofile.Profile {
	if c.Theme == theme.Mono {
		return colorprofile.ASCII
	}
	switch c.Profile {
	case "ansi256":
		return colorprofile.ANSI256
	case "ansi16":
		return colorprofile.ANSI
	case "none":
		return colorprofile.ASCII
	}
	return colorprofile.TrueColor
}

type Environment struct {
	Config Config
	Emit   func(any)
}

type Definition struct {
	ID          string                                   `json:"id"`
	Title       string                                   `json:"title"`
	Level       string                                   `json:"level"`
	Family      string                                   `json:"family"`
	Description string                                   `json:"description"`
	Source      string                                   `json:"source"`
	Symbol      string                                   `json:"symbol"`
	Upstream    string                                   `json:"upstream"`
	Controls    []string                                 `json:"controls"`
	States      []string                                 `json:"states"`
	Scenarios   []string                                 `json:"scenarios"`
	Clock       bool                                     `json:"clock"`
	Step        bool                                     `json:"step"`
	Defaults    Config                                   `json:"defaults"`
	Factory     func(Environment) (runtime.Scene, error) `json:"-"`
}

// Definitions returns fresh metadata; factories always construct fresh model state.
func Definitions() []Definition { return definitions() }
func Find(id string) (Definition, error) {
	for _, d := range definitions() {
		if d.ID == id {
			return d, nil
		}
	}
	return Definition{}, fmt.Errorf("unknown story %q", id)
}
func New(id string, config Config, emit func(any)) (*runtime.Root, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	d, err := Find(id)
	if err != nil {
		return nil, err
	}
	scene, err := d.Factory(Environment{Config: config, Emit: emit})
	if err != nil {
		return nil, err
	}
	scene.Resize(config.Size())
	return &runtime.Root{Scene: scene, Background: config.Palette().Background, Size: config.Size(), AltScreen: true, ASCII: config.Palette().ASCII}, nil
}

// Frame executes the real Go View without starting I/O or wall-clock commands.
func Frame(id string, config Config) (string, error) {
	m, err := New(id, config, nil)
	if err != nil {
		return "", err
	}
	defer m.Close()
	return FormatFrame(m.View().Content, config), nil
}
func FormatFrame(frame string, config Config) string {
	var out bytes.Buffer
	w := colorprofile.Writer{Forward: &out, Profile: config.ColorProfile()}
	_, _ = w.WriteString(frame)
	return out.String()
}
