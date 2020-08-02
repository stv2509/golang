package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ReadDir(dir string) (map[string]string, error)
var dirfile = "C:\\Users\\U_M12QS\\Desktop\\download\\golang\\HW-11-13\\ENV\\"

func ReadDir(dirfile string) (map[string]string, error) {

	cache := make(map[string]string)

	dir, err := os.Open(dirfile)

	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdirnames(-1)

	if err != nil {
		return nil, err
	}

	for _, envname := range fileInfos {

		value, err := ioutil.ReadFile(dirfile + envname)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cache[envname] = strings.TrimSpace(string(value))
	}
	return cache, nil

}
