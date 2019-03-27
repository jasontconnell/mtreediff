package process

import (
	"sort"

	"github.com/jasontconnell/mtreediff/data"
)

func Compare(folders []*data.Folder) []*data.CompareResult {
	list := []*data.CompareResult{}
	for i, f := range folders {
		for j, fj := range folders {
			if i == j {
				continue
			}

			list = append(list, &data.CompareResult{Folders: []*data.Folder{f, fj}})
		}
	}

	doCompare(list)
	return list
}

func CompareAll(folders []*data.Folder) []*data.CompareResult {
	cr := &data.CompareResult{Folders: folders}
	list := []*data.CompareResult{cr}
	doCompare(list)
	return list
}

func doCompare(list []*data.CompareResult) {
	for _, c := range list {
		maps := []map[string]*data.File{}
		for _, folder := range c.Folders {

			m := make(map[string]*data.File)

			for _, f := range folder.Files {
				m[f.RelPath] = f
			}

			maps = append(maps, m)
		}

		base := maps[0]
		for fn, f := range base {
			same, diff, miss := 0, 0, 0
			for _, m := range maps[1:] {
				rf, ok := m[fn]

				if ok {
					if f.SHA == rf.SHA {
						same++
					} else {
						diff++
					}
				} else {
					miss++
				}
			}
			if same == len(maps[1:]) {
				c.Same = append(c.Same, f)
			} else if diff > 0 {
				c.Diff = append(c.Diff, f)
			} else if miss > 0 {
				c.Miss = append(c.Miss, f)
			}
		}

		sort.Slice(c.Diff, func(i, j int) bool {
			return c.Diff[i].RelPath < c.Diff[j].RelPath
		})
		sort.Slice(c.Same, func(i, j int) bool {
			return c.Same[i].RelPath < c.Same[j].RelPath
		})
		sort.Slice(c.Miss, func(i, j int) bool {
			return c.Miss[i].RelPath < c.Miss[j].RelPath
		})
	}
}
