package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestWorkbenchCommandPresentation(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	for _, tc := range []struct {
		args []string
		want string
	}{
		{[]string{"--help"}, "Examples\n  workbench run screen/services"},
		{[]string{"help", "capture"}, "--columns int"},
		{[]string{"--version"}, "workbench version 0.1.0"},
		{[]string{"man"}, ".TH WORKBENCH 1"},
		{[]string{"completion", "bash"}, "__start_workbench"},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			var output bytes.Buffer

			root := command()
			root.SetOut(&output)
			root.SetArgs(tc.args)

			if err := root.ExecuteContext(t.Context()); err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(output.String(), tc.want) {
				t.Errorf("missing %q in output:\n%s", tc.want, output.String())
			}
		})
	}
}
