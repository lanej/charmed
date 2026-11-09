package workbench

import (
	"errors"
	"os"
)

func writeGuideFile(directory, name string, data []byte) (err error) {
	root, err := os.OpenRoot(directory)
	if err != nil {
		return err
	}

	defer func() { err = errors.Join(err, root.Close()) }()

	return root.WriteFile(name, data, 0o644)
}
