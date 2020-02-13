package openfile

import (
	"errors"
	"io/ioutil"
	"os"
)

//OpenFile - open file
func OpenFile() (string, error) {

	var str string
	var err error

	if len(os.Args) > 1 {
		file := os.Args[1]

		bs, err := ioutil.ReadFile(file)

		if err != nil {
			err := errors.New("\nCan`t open file \"" + file + "\"")
			return str, err
		}

		str := string(bs)
		return str, err

	}

	err = errors.New("\nInput file not found. Select a file")
	return str, err

}
