package sll

import (
	"errors"
	"fmt"
)

var ErrIndexIsOutOfSize = errors.New("index is out of size")

type node[T any] struct {
	next *node[T]
	val  T
}

type SLList[T any] struct {
	head *node[T]
	size int
}

func (l *SLList[T]) getNodeByIdx(idx int) *node[T] {
	current := l.head

	for count := 0; count < idx; count++ {
		current = current.next
	}

	return current
}

func (l *SLList[T]) Insert(elem T) {
	node := &node[T]{val: elem}

	if l.head == nil {
		l.head = node
		l.size = 1
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	l.size++
	current.next = node
}

func (l *SLList[T]) Traverse(f func(v any)) {
	current := l.head

	for current != nil {
		f(current.val)
		current = current.next
	}
}

func (l *SLList[T]) IsEmpty() bool {
	if l.head == nil {
		return true
	} else {
		return false
	}
}

func (l *SLList[T]) Size() int {
	return l.size
}

func (l *SLList[T]) At(idx int) (T, error) {
	if l.IsEmpty() || idx < 0 || idx >= l.size {
		var tNil T
		return tNil, ErrIndexIsOutOfSize
	}

	current := l.getNodeByIdx(idx)

	return current.val, nil
}

func (l *SLList[T]) DeleteAt(idx int) error {
	if idx < 0 || idx >= l.Size() {
		return ErrIndexIsOutOfSize
	}

	l.size--

	if idx == 0 {
		l.head.val = l.head.next.val
		l.head.next = l.head.next.next
		return nil
	}

	current := l.getNodeByIdx(idx - 1)

	current.next = current.next.next
	return nil
}

func (l *SLList[T]) InsertFront(t T) {
	node := &node[T]{val: t}
	node.next = l.head
	l.head = node
	l.size++
}

// InsertAt Вставка элемента на позицию idx
func (l *SLList[T]) InsertAt(idx int, t T) error {
	if idx < 0 {
		return ErrIndexIsOutOfSize
	}

	if idx >= l.Size() {
		fmt.Println("the index is bigger than list Size. value will be inserted at the end of the list")
		l.Insert(t)
		return nil
	}

	if idx == 0 {
		l.InsertFront(t)
		return nil
	}

	nd := &node[T]{val: t}

	current := l.getNodeByIdx(idx - 1)

	nd.next = current.next
	current.next = nd
	l.size++
	return nil
}

//func (l *SLList[T]) sort() {
//	if l.IsEmpty() {
//		fmt.Println("the list is empty")
//		return
//	}
//
//	x := l.Size()
//	data := make([]T, 0, x)
//
//	current := l.head
//	count := 0
//
//	for current != nil {
//		data[count] = current.val
//		count++
//		current = current.next
//	}
//
//	for i := 0; i < (x - 1); i++ {
//		for j := (x - 1); j > i; j-- {
//			if data[j-1] > data[j] {
//				temp := data[j-1]
//				data[j-1] = data[j]
//				data[j] = temp
//			}
//		}
//	}
//
//	current = l.head
//	count = 0
//
//	for current != nil {
//		current.val = data[count]
//		count++
//		current = current.next
//	}
//}
