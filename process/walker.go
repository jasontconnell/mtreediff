package process

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jasontconnell/mtreediff/data"
)

func Walk(dir string) (*data.Folder, error) {
	stat, err := os.Stat(dir)
	if stat != nil && !stat.IsDir() {
		err = errors.New("not a directory")
	}

	if err != nil {
		return nil, err
	}

	var full string

	if !filepath.IsAbs(dir) {
		full, err = filepath.Abs(stat.Name())
		if err != nil {
			return nil, err
		}
	} else {
		full = dir
	}

	full = strings.TrimSuffix(full, "\\")
	list := []*data.File{}

	filepath.Walk(full, func(path string, info os.FileInfo, err error) error {
		fulldir, _ := filepath.Split(path)
		fulldir = strings.TrimSuffix(fulldir, "\\")
		relpath := strings.TrimPrefix(path, dir)

		name := info.Name()
		if !info.IsDir() {
			f := &data.File{Name: name, RelPath: relpath, FullPath: path}
			list = append(list, f)
		}

		return nil
	})

	m := make(map[string]*data.File)
	for _, f := range list {
		m[f.RelPath] = f
	}

	_, name := filepath.Split(dir)
	folder := &data.Folder{Path: dir, Name: name, Files: m}

	return folder, nil
}

func GetSubdirs(base string) ([]string, error) {
	if !filepath.IsAbs(base) {
		base, _ = os.Getwd()
	}

	folders, err := ioutil.ReadDir(base)
	if err != nil {
		return nil, err
	}

	list := []string{}
	for _, f := range folders {
		if f.IsDir() {
			list = append(list, filepath.Join(base, f.Name()))
		}
	}
	sort.Strings(list)
	return list, nil
}

func GetDirs(base, str string) []string {
	dirs := strings.Split(str, ",")
	list := []string{}

	for _, d := range dirs {
		list = append(list, filepath.Join(base, d))
	}
	return list
}
