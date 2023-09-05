package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Head *ListItem
	Tail *ListItem
}

func (l *list) Len() int {
	length := 0
	current := l.Head

	for current != nil {
		length++
		current = current.Next
	}

	return length
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Next = l.Head
		l.Head.Prev = newItem
		l.Head = newItem

	}
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.Head == nil {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Prev = l.Tail
		l.Tail.Next = newItem
		l.Tail = newItem
	}

	return newItem
}

func (l *list) Remove(i *ListItem) {
	if i == l.Head {
		l.Head = i.Next
	}

	if i == l.Tail {
		l.Tail = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.Head {
		return
	}

	if i == l.Tail {
		l.Tail = i.Prev
		l.Tail.Next = nil
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.Head
	l.Head.Prev = i
	l.Head = i
}

func NewList() List {
	return new(list)
}
