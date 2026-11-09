package story

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func permissionsStory(e Environment) (runtime.Scene, error) {
	return static(func(layout.Size) string {
		t := e.Config.Palette()
		rows := []string{
			"Permission combinations / one group",
			"Read: top left. Write: top right.",
			"Execute: bottom left.",
			"",
		}

		for _, example := range []struct {
			mode fs.FileMode
			name string
		}{
			{0o400, "read"},
			{0o600, "read + write"},
			{0o500, "read + execute"},
			{0o700, "read + write + execute"},
			{0o200, "write"},
			{0o100, "execute"},
			{0o300, "write + execute"},
			{0, "none"},
		} {
			p := text.Permissions{Mode: example.mode, Available: true}
			groups := p.Groups(t)
			rows = append(rows, groups[0]+"  "+p.Label()[:3]+"  "+example.name)
		}

		rows = append(rows, "", "Owner / group / other    exact mode")

		for _, example := range []struct {
			mode fs.FileMode
			name string
		}{
			{0o755, "docs/"},
			{0o644, "app.ts"},
			{0o755, "deploy.sh"},
			{0o600, ".env"},
			{0o444, "report.pdf"},
		} {
			p := text.Permissions{Mode: example.mode, Available: true}
			rows = append(rows, fmt.Sprintf("%s  %-10s %s", p.Compact(t), example.name, p.Label()))
		}

		unknown := text.Permissions{}
		rows = append(rows, unknown.Compact(t)+"  uninspected "+unknown.Label(), "", "Special bits / explicit modifiers")

		for _, mode := range []fs.FileMode{0o755 | fs.ModeSetuid, 0o644 | fs.ModeSetgid, 0o777 | fs.ModeSticky} {
			p := text.Permissions{Mode: mode, Available: true}
			rows = append(rows, p.Compact(t)+"  "+p.Label())
		}

		rows = append(rows, "", "Mode bits are not effective access.")

		return strings.Join(rows, "\n")
	}), nil
}
