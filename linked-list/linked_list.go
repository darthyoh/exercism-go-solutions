package linkedlist

import "errors"

var ErrEmptyList = errors.New("Empty List")

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	tail *Node
	size int
}

func NewList(args ...interface{}) *List {
	l := List{}
	if len(args) == 0 {
		return &l
	}
	for _, e := range args {
		l.PushBack(e)
	}

	return &l
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) PushBack(elem interface{}) {
	newNode := Node{elem, nil, nil}
	l.tail = &newNode
	if l.size == 0 {
		l.head = &newNode

	} else {
		active := l.head
		for {
			if active.next == nil {
				newNode.prev = active
				active.next = &newNode

				break
			}
			active = active.next
		}
	}
	l.size++
}

func (l *List) PushFront(elem interface{}) {
	newNode := Node{elem, nil, nil}
	if l.size == 0 {
		l.head = &newNode
		l.tail = &newNode
	} else {
		newNode.next = l.head
		l.head.prev = &newNode
		l.head = &newNode
	}
	l.size++
}

func (l *List) PopBack() (lastElem interface{}, error error) {
	lastElem = nil
	error = nil

	switch l.size {
	case 0:
		error = ErrEmptyList
	case 1:
		lastElem = l.tail.Val
		l.tail = nil
		l.head = nil
		l.size = 0
	default:
		lastElem = l.tail.Val
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
	}
	return
}

func (l *List) PopFront() (firstElem interface{}, error error) {
	firstElem = nil
	error = nil

	switch l.size {
	case 0:
		error = ErrEmptyList
	case 1:
		firstElem = l.head.Val
		l.head = nil
		l.tail = nil
		l.size = 0
	default:
		firstElem = l.head.Val
		l.head = l.head.next
		l.head.prev = nil
		l.size--
	}
	return
}

func (l *List) Reverse() *List {
	if l.size == 0 {
		return l
	}

	newList := List{nil, nil, 0}
	active := l.tail

	for {
		newList.PushBack(active.Val)
		if active.prev == nil {
			break
		}
		active = active.prev
	}
	l.head = newList.head
	l.tail = newList.tail
	return l
}
