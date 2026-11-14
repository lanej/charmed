package workbench

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var guideSkills = []string{"charmed", "charmed-contributing"}

func writeGuideSkills(directory, root string) error {
	for _, name := range guideSkills {
		source := os.DirFS(filepath.Join(root, ".agents", "skills", name))

		var output bytes.Buffer

		compressed := gzip.NewWriter(&output)
		archive := tar.NewWriter(compressed)

		err := fs.WalkDir(source, ".", func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}

			if entry.IsDir() {
				if entry.Name() == "__pycache__" || (path != "." && strings.HasPrefix(entry.Name(), ".")) {
					return fs.SkipDir
				}

				return nil
			}

			if strings.HasPrefix(entry.Name(), ".") {
				return nil
			}

			if !entry.Type().IsRegular() {
				return fmt.Errorf("skill %s: %s must be a regular file", name, path)
			}

			switch filepath.Ext(path) {
			case ".md", ".py", ".sh", ".yml", ".json":
			default:
				return fmt.Errorf("skill %s: unsupported package file %s", name, path)
			}

			data, readErr := fs.ReadFile(source, path)
			if readErr != nil {
				return readErr
			}

			target := filepath.Join("skills", name, path)
			if mkdirErr := os.MkdirAll(filepath.Dir(filepath.Join(directory, target)), 0o755); mkdirErr != nil { //nolint:gosec // WHY: Published directories must be traversable by the static server.
				return mkdirErr
			}

			if writeErr := writeGuideFile(directory, filepath.ToSlash(target), data); writeErr != nil {
				return writeErr
			}

			header := &tar.Header{Name: name + "/" + path, Mode: 0o644, Size: int64(len(data)), Typeflag: tar.TypeReg}
			if headerErr := archive.WriteHeader(header); headerErr != nil {
				return headerErr
			}

			_, writeErr := archive.Write(data)

			return writeErr
		})
		if err = errors.Join(err, archive.Close(), compressed.Close()); err != nil {
			return fmt.Errorf("publish skill %s: %w", name, err)
		}

		if _, statErr := fs.Stat(source, "SKILL.md"); statErr != nil {
			return fmt.Errorf("skill %s entrypoint: %w", name, statErr)
		}

		if writeErr := writeGuideFile(directory, "skills/"+name+".tar.gz", output.Bytes()); writeErr != nil {
			return writeErr
		}
	}

	return nil
}
