package main

import (
	"fmt"
	"github.com/stv2509/golang/HW-2/unpacker"
)

func main() {

	var str string

	fmt.Print("\nВведите строку: ")
	fmt.Scanln(&str)

	res, err := unpacker.Unpack(str)

	if err != nil {

		fmt.Println("String Error: ", err)

	} else {

		fmt.Println("Result: ", res)
	}
}
