package process

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jasontconnell/mtreediff/data"
)

func ReadAll(folder *data.Folder) {
	for _, f := range folder.Files {
		ff, err := os.Open(f.FullPath)
		if err != nil {
			log.Printf("couldn't open file %s.\n", f.FullPath)
			ff.Close()
			continue
		}

		sum, err := sha256sum(ff)
		if err != nil {
			log.Printf("couldn't read file %s.\n", f.FullPath)
			ff.Close()
			continue
		}
		sumstr := hex.EncodeToString(sum)

		f.SHA = sumstr
		ff.Close()
	}
}

func sha256sum(r io.Reader) ([]byte, error) {
	h := sha256.New()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	h.Write(b)
	return h.Sum(nil), nil
}
