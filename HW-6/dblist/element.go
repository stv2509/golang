package dblist

//Item - элемент списка
type Item struct {
	pNext *Item
	pPrev *Item
	value interface{}
}

//Value - возвращает значение
func (i *Item) Value() interface{} {

	return i.value
}

//Next - следующий Item
func (i *Item) Next() *Item {

	return i.pNext
}

//Prev - предыдущий Item
func (i *Item) Prev() *Item {

	return i.pPrev
}
