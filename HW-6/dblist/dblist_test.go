package dblist

import (
	//	"bytes"
	//	"reflect"
	"fmt"
	"testing"
)

var (
	dl, dlt                         *List
	itemFirst, itemLast, itemRemove *Item
)

func init() {

	dl = NewList()
	dlt = NewList()
	itemFirst = dlt.PushFront("Elem 1")
	dlt.PushBack("Elem 2")
	itemRemove = dlt.PushBack("Elem 3")
	itemLast = dlt.PushBack("Elem 4")

}

func TestLen(t *testing.T) {
	got := dlt.Len()
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFirst(t *testing.T) {

	got, _ := dlt.First()

	want := itemFirst

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestLast(t *testing.T) {

	got, _ := dlt.Last()

	want := itemLast

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestFirstZ(t *testing.T) {

	_, err := dl.First()

	if err != errList {
		t.Errorf("Last() error = %v, wantErr %v", err, errList)
	}

}

func TestLastZ(t *testing.T) {

	_, err := dl.Last()

	if err != errList {
		t.Errorf("Last() error = %v, wantErr %v", err, errList)
	}

}

func Example() {

	list := NewList()
	list.PushFront("Item PushFront 1")
	list.PushBack("Item PushBack")
	list.PushFront("Item PushFront 2")
	itA := list.PushBack("Item A")

	fmt.Println("ShowList:")
	for e, _ := list.First(); e != nil; e = e.Next() {

		fmt.Println(e.Value())
	}
	fmt.Println("Count item:", list.Len())
	fmt.Println("Remove Item A")
	fmt.Println("ShowList:")
	list.Remove(itA)
	for e, _ := list.Last(); e != nil; e = e.Prev() {

		fmt.Println(e.Value())
	}
	fmt.Println("Count item:", list.Len())
	// Output:
	// ShowList:
	// Item PushFront 2
	// Item PushFront 1
	// Item PushBack
	// Item A
	// Count item: 4
	// Remove Item A
	// ShowList:
	// Item PushBack
	// Item PushFront 1
	// Item PushFront 2
	// Count item: 3
}
