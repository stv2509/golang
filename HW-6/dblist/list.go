package dblist

import (
	"errors"
)

var errList = errors.New("Empty list")

//List - тип контейнер
type List struct {
	size      int
	firstItem *Item
	lastItem  *Item
}

//NewList - initializet new dblist
func NewList() *List {

	return &List{}
}

//Len - длинна списка
func (l *List) Len() int {

	return l.size
}

//First - первый Item
func (l *List) First() (*Item, error) {

	if l.firstItem == nil {

		return nil, errList
	}
	return l.firstItem, nil
}

//Last - последний Item
func (l *List) Last() (*Item, error) {

	if l.lastItem == nil {

		return nil, errList
	}
	return l.lastItem, nil
}

//PushFront - добавить значение в начало
func (l *List) PushFront(v interface{}) *Item {

	newItem := &Item{value: v}

	if val := l.firstItem; val == nil {

		l.firstItem, l.lastItem = newItem, newItem

	} else {

		newItem.pNext = l.firstItem

		l.firstItem.pPrev = newItem

		l.firstItem = newItem
	}

	l.size++

	return newItem
}

//PushBack - добавить значение в конец
func (l *List) PushBack(v interface{}) *Item {

	newItem := &Item{value: v}

	if val := l.lastItem; val == nil {

		l.firstItem, l.lastItem = newItem, newItem

	} else {

		newItem.pPrev = l.lastItem

		l.lastItem.pNext = newItem

		l.lastItem = newItem
	}

	l.size++

	return newItem
}

//Remove - удалить элемент
func (l *List) Remove(i *Item) {

	if i.pPrev == nil {

		l.firstItem = i.pNext

	} else {

		i.pPrev.pNext = i.pNext

	}
	if i.pNext == nil {

		l.lastItem = i.pPrev

	} else {

		i.pNext.pPrev = i.pPrev
	}

	l.size--
}
