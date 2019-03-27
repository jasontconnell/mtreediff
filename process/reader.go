package process

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jasontconnell/mtreediff/data"
)

func ReadAll(folder *data.Folder) {
	for _, f := range folder.Files {
		ff, err := os.Open(f.FullPath)
		if err != nil {
			fmt.Println("couldn't open ", f.FullPath)
			continue
		}

		sum := sha256sum(ff)
		sumstr := hex.EncodeToString(sum)

		f.SHA = sumstr
	}
}

func sha256sum(r io.Reader) []byte {
	h := sha256.New()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println("couldn't read")
	}
	h.Write(b)
	return h.Sum(nil)
}
