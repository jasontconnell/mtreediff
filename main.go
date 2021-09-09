package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jasontconnell/mtreediff/data"
	"github.com/jasontconnell/mtreediff/process"
)

func main() {
	start := time.Now()
	dirs := flag.String("d", "", "compare the directories specified")
	subdirs := flag.String("s", "", "compare the subdirectories below specified directory")
	out := flag.String("o", "", "output folder")
	rm := flag.Bool("rm", false, "clear output before compare?")
	rev := flag.Bool("reverse", false, "reverse folders (makes it easier to just add --reverse to existing call in the console)")
	flag.Parse()

	if (dirs == nil || *dirs == "") && (subdirs == nil || *subdirs == "") {
		flag.PrintDefaults()
		log.Fatal("dirs (-d) or subdirs (-s) is required")
	}

	if out == nil || *out == "" {
		flag.PrintDefaults()
		log.Fatal("output folder (-o) required")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("couldn't get current directory", err)
	}

	var folders []string
	if *dirs != "" {
		folders = process.GetDirs(wd, *dirs, *rev)
	} else if *subdirs != "" {
		var err error
		folders, err = process.GetSubdirs(*subdirs)
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

	if *rm {
		err := os.RemoveAll(*out)

		if err != nil {
			log.Fatalf("clearing directory before output %s. %v", *out, err)
		}
	}

	results := process.CompareAll(allmaps)

	for _, c := range results {
		name := ""
		for _, f := range c.Folders {
			name += f.Name + "-"
		}
		outpath := filepath.Join(*out, strings.TrimRight(name, "-"))
		diffpath := filepath.Join(outpath, "diff")
		process.Copy(diffpath, c.Diff)

		misspath := filepath.Join(outpath, "miss")
		process.Copy(misspath, c.Miss)

		samepath := filepath.Join(outpath, "same")
		process.Copy(samepath, c.Same)
	}

	fmt.Println("Finished", time.Since(start))
}
