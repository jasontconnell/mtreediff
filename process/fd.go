package process

import (
	"io"
	"os"
	"path/filepath"

	"github.com/jasontconnell/mtreediff/data"
	"github.com/pkg/errors"
)

func Copy(dest string, list []*data.File) error {
	for _, f := range list {
		path := filepath.Join(dest, f.RelPath)
		folder, _ := filepath.Split(path)
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "Couldn't make dir %s", folder)
		}

		r, err := os.Open(f.FullPath)
		if err != nil {
			return errors.Wrapf(err, "Couldn't open file %s", f.FullPath)
		}

		w, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		if err != nil {
			return errors.Wrapf(err, "Couldn't open file for copying %s", path)
		}

		_, err = io.Copy(w, r)
		if err != nil {
			return errors.Wrapf(err, "Couldn't copy file %s", folder)
		}
	}

	return nil
}
