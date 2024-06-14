package dll

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	prev *node[T]
	next *node[T]
	val  T
}

type DLList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

var ErrIndexIsOutOfSize = errors.New("index is out of size")

func (l *DLList[T]) getNodeByIdx(idx int) *node[T] {
	current := l.head

	for count := 0; count < idx; count++ {
		current = current.next
	}

	return current
}

func (l *DLList[T]) Insert(elem T) {
	node := &node[T]{val: elem}

	l.size++

	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node
}

func (l *DLList[T]) GetTail() (T, error) {
	if l.tail == nil {
		var tNil T
		return tNil, ErrIndexIsOutOfSize
	}
	return l.tail.val, nil
}

func (l *DLList[T]) Reverse() {
	if l.Size() < 2 {
		return
	}

	current := l.head
	var temp *node[T]

	l.head, l.tail = l.tail, l.head

	for current != nil {
		temp = current.next
		current.prev, current.next = current.next, current.prev
		current = temp
	}
}

func (l *DLList[T]) Traverse(f func(v any)) {
	current := l.head

	for current != nil {
		f(current.val)
		current = current.next
	}
}

func (l *DLList[T]) IsEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *DLList[T]) Size() int {
	return l.size
}

func (l *DLList[T]) At(idx int) (T, error) {
	if l.IsEmpty() || idx < 0 || idx >= l.size {
		var tNil T
		return tNil, ErrIndexIsOutOfSize
	}

	current := l.getNodeByIdx(idx)

	return current.val, nil
}

func (l *DLList[T]) DeleteAt(idx int) error {
	if l.IsEmpty() || idx < 0 || idx >= l.size {
		return ErrIndexIsOutOfSize
	}

	if idx == 0 && l.size == 1 {
		l.head, l.tail, l.size = nil, nil, 0
		return nil
	}

	l.size--

	if idx == 0 {
		l.head.val = l.head.next.val
		l.head.next = l.head.next.next
		l.head.prev = nil
		return nil
	}

	if idx == l.size-1 {
		if err := l.DeleteFromTail(); err != nil {
			return fmt.Errorf("delete from tail: %w", err)
		}
		return nil
	}

	current := l.getNodeByIdx(idx - 1)

	current.next, current.next.next.prev = current.next.next, current

	return nil
}

func (l *DLList[T]) DeleteFromTail() error {
	if l.IsEmpty() {
		return ErrIndexIsOutOfSize
	}

	if l.size == 1 {
		l.head, l.tail, l.size = nil, nil, 0
		return nil
	}

	if l.size == 2 {
		l.tail = l.head
		l.head.next = nil
		l.size = 1
		return nil
	}

	l.size--

	current := l.head
	for current.next.next != nil {
		current = current.next
	}
	l.tail = current

	return nil
}

func (l *DLList[T]) InsertAt(idx int, t T) error {
	if idx < 0 || idx >= l.size {
		return ErrIndexIsOutOfSize
	}

	if idx == 0 {
		l.InsertFront(t)
		return nil
	}

	node := &node[T]{val: t}

	current := l.getNodeByIdx(idx - 1)

	node.next, node.prev = current.next, current
	current.next, current.next.prev = node, node
	l.size++
	return nil
}

func (l *DLList[T]) InsertFront(t T) {
	if l.head == nil {
		l.Insert(t)
		return
	}
	node := &node[T]{val: t}
	node.next, l.head.prev = l.head, node
	l.head = node
	l.size++
}
