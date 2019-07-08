package source

import (
	"io/ioutil"
	"log"
	"os"
)

func mustReadFile(path string) []byte {
	fd, err := os.Open(path)
	if err != nil {
		msg := "failed to resolve file"
		if err == os.ErrNotExist {
			msg = "file does not exist"
		}

		log.Fatalf("Error resolving file `%s`: %v %v\n", path, msg, err)
	}

	defer fd.Close()
	b, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatalf("Error reading file `%s`: %s\n", path, err.Error())
	}

	return b
}
