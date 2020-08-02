package main

import (
	"fmt"
)

var allW = 120
var N = 37
var n = N
var j = 0

func main() {

	for j < allW {

		fmt.Printf("\n%v : %v gorutiens\n", j, N)
		j = N
		if n+j < allW {
			N = n+j
		} else {
			N = allW
		}
	}

}
