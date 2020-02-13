package main

import (
	"fmt"
	"HW-4/sortmax"
)

var pf = fmt.Printf

func main() {

	s := []interface{}{10, 30, 55, 1341, 128, 576}

	max, err := sortmax.GetMaxValue(s, sortmax.GetComparator(s))

	pf("\nResult: %v, %v", max, err)


}