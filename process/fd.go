package process

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jasontconnell/mtreediff/data"
)

func Copy(dest string, list []*data.File) error {
	for _, f := range list {
		path := filepath.Join(dest, f.RelPath)
		folder, _ := filepath.Split(path)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't make dir %s. %w", folder, err)
		}

		r, err := os.Open(f.FullPath)
		if err != nil {
			return fmt.Errorf("couldn't open file %s", f.FullPath)
		}

		w, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't open file for copying %s", path)
		}

		_, err = io.Copy(w, r)
		if err != nil {
			return fmt.Errorf("couldn't copy file %s", folder)
		}
	}

	return nil
}
