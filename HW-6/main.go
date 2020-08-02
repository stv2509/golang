package main

import (
	l "HW-6/dblist"
	"fmt"
	"strings"
)

func dashes() {
	dashes := strings.Repeat("-", 50)
	fmt.Println(dashes)
	return
}

func main() {
	dashes()
	dll := l.NewList()
	dll.PushFront(789)
	itB := dll.PushFront("itB")
	dll.PushBack(75)
	dll.PushBack([]interface{}{324, 123, 78})
	dll.PushFront([]interface{}{"root", "toor", "linux"})

	for e, _ := dll.First(); e != nil; e = e.Next() {

		fmt.Println(e.Value())
		dashes()
	}

	dll.Remove(itB)
	
	fmt.Print("\ndll.Remove(itB)\n\n")
	for e, _ := dll.First(); e != nil; e = e.Next() {

		fmt.Println(e.Value())
		dashes()
	}
	
	fmt.Println(dll.Len())
}
