package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/jasontconnell/mtreediff/data"
	"github.com/jasontconnell/mtreediff/process"
)

func main() {
	dirs := flag.String("d", "", "compare the directories specified")
	subdirs := flag.String("s", "", "compare the subdirectories below specified directory specified")
	out := flag.String("o", "", "output folder")
	flag.Parse()

	if (dirs == nil || *dirs == "") && (subdirs == nil || *subdirs == "") {
		flag.PrintDefaults()
		log.Fatal("dirs (-d) or subdirs (-s) is required")
	}

	if out == nil || *out == "" {
		flag.PrintDefaults()
		log.Fatal("output folder (-o) required")
	}

	var folders []string
	if *dirs != "" {
		folders = strings.Split(*dirs, ",")
	} else if *subdirs != "" {
		var err error
		folders, err = process.GetDirs(*subdirs)
		if err != nil {
			log.Fatalf("Couldn't load subdirs %v %v", *subdirs, err)
		}
	}

	allmaps := []*data.Folder{}

	for _, f := range folders {
		r, err := process.Walk(f)
		if err != nil {
			log.Fatal(err)
		}

		process.ReadAll(r)
		allmaps = append(allmaps, r)
	}

	results := process.CompareAll(allmaps)
	fmt.Println("All maps len", len(allmaps))

	for _, c := range results {
		name := ""
		for _, f := range c.Folders {
			name += f.Name + "-"
		}
		outpath := filepath.Join(*out, strings.TrimRight(name, "-"))
		// outfile, err := os.OpenFile(outpath+".txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
		// if err != nil {
		// 	fmt.Println("couldn't open output file", outpath+".txt")
		// 	continue
		// }
		// defer outfile.Close()

		// for _, f := range c.Diff {
		diffpath := filepath.Join(outpath, "diff")
		//fmt.Fprintf(outfile, "Diff %v\n", f.RelPath)
		process.Copy(diffpath, c.Diff)
		// }

		misspath := filepath.Join(outpath, "miss")
		process.Copy(misspath, c.Miss)

		samepath := filepath.Join(outpath, "same")
		process.Copy(samepath, c.Same)

		// for _, f := range c.Miss {
		// 	fmt.Fprintf(outfile, "Missed %v\n", f.RelPath)
		// }

		// for _, f := range c.Same {
		// 	fmt.Fprintf(outfile, "Same %v\n", f.RelPath)
		// }
	}
}
