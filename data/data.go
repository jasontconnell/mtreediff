package data

import (
	"fmt"
)

type File struct {
	Name     string
	RelPath  string
	FullPath string
	SHA      string
}

type Folder struct {
	Path  string
	Name  string
	Files map[string]*File
}

func (f File) String() string {
	return fmt.Sprintf("%s -- %s", f.RelPath, f.Name)
}

type CompareResult struct {
	Folders []*Folder
	Same    []*File
	Diff    []*File
	Miss    []*File
}
