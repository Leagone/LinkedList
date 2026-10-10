package main

import (
	"errors"
)

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func toSlice(l *LinkedList) []int {
	var result []int
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (l *LinkedList) Append(value int) {

	newNode := &Node{
		Value: value,
	}

	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
		newNode.Prev = current
	}

	l.Tail = newNode
	l.Length++

}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{
		Value: value,
	}
	newNode.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) Pop() (int, error) {

	if l.Head == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	val := l.Tail.Value

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Tail.Prev.Next = nil
		l.Tail.Prev = nil
	}

	l.Length--
	return val, nil

}

func (l *LinkedList) Get(index int) (int, error) {

	if index < 0 {
		return 0, errors.New("Out of bounds")
	}

	current := l.Head
	for i := 0; current != nil; i++ {
		if i == index {
			return current.Value, nil
		}
		current = current.Next
	}

	return 0, errors.New("Out of Bounds")
}

func (l *LinkedList) InsertAt(index int, value int) error {
	if index < 0 {
		return errors.New("index out of bounds")
	}

	newNode := &Node{Value: value}

	if index == 0 {
		if l.Head == nil {
			l.Head = newNode
			l.Tail = newNode
		} else {
			newNode.Next = l.Head
			l.Head.Prev = newNode
			l.Head = newNode
		}
		l.Length++
		return nil
	}

	current := l.Head
	for i := 0; i < index; i++ {
		if current == nil {
			return errors.New("index out of bounds")
		}
		current = current.Next
	}

	if current == nil {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
		l.Length++
		return nil
	}

	newNode.Next = current
	newNode.Prev = current.Prev

	current.Prev.Next = newNode
	current.Prev = newNode

	l.Length++
	return nil
}

func (l *LinkedList) RemoveAt(index int) error {
	if index < 0 || l.Head == nil {
		return errors.New("Out of Bounds")
	}

	if index == 0 {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}
		l.Length--
		return nil
	}

	current := l.Head

	for i := 0; i < index; i++ {
		current = current.Next
		if current == nil {
			return errors.New("Out of Bounds")
		}
	}

	if current.Next == nil {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		l.Length--
		return nil
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	l.Length--
	return nil

}

func (l *LinkedList) RemoveByValue(value int) error {

	if l.Head == nil {
		return errors.New("Cannot remove from empty list")
	}

	if l.Head.Value == value {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}
		l.Length--
		return nil

	}

	current := l.Head
	for current != nil {
		current = current.Next
		if current == nil {
			return errors.New("Value not found")
		}
		if current.Value == value {
			break
		}
	}

	if current.Next == nil {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		l.Length--
		return nil
	}

	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	l.Length--
	return nil

}

func (l *LinkedList) Contains(value int) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) IndexOf(value int) (int, error) {
	current := l.Head
	for i := 0; current != nil; i++ {
		if current.Value == value {
			return i, nil
		}
		current = current.Next
	}

	return -1, errors.New("value not found")
}

func (l *LinkedList) Filter(condition func(value int) bool) {
	dummy := &Node{}

	tail := dummy
	newLenght := 0
	current := l.Head

	for current != nil {
		next := current.Next
		if condition(current.Value) {
			tail.Next = current
			current.Prev = tail
			tail = current
			newLenght++
		}
		current = next
	}

	tail.Next = nil

	l.Head = dummy.Next
	l.Length = newLenght

	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Prev = nil
		l.Tail = tail
	}
}

func (l *LinkedList) Reverse() {
	current := l.Head
	for current != nil {
		current.Next, current.Prev = current.Prev, current.Next
		current = current.Prev
	}

	l.Head, l.Tail = l.Tail, l.Head
}

func (l *LinkedList) Middle() int {
	if l.Head == nil {
		return -1
	}

	slow := l.Head
	fast := l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Value
}

func (l *LinkedList) RemoveNthFromEnd(n int) error {

	if n <= 0 || l.Head == nil || n > l.Length {
		return errors.New("Out of Bounds")
	}

	current := l.Tail
	for i := 1; i < n; i++ {
		current = current.Prev
	}

	switch current {
	case l.Head:
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		} else {
			l.Tail = nil
		}
	case l.Tail:
		l.Tail = l.Tail.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		} else {
			l.Head = nil
		}
	default:
		current.Next.Prev = current.Prev
		current.Prev.Next = current.Next
	}

	l.Length--
	return nil
}

func (l *LinkedList) HasCycle() bool {
	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func (l *LinkedList) FindCycleStart() *Node {
	slow := l.Head
	fast := l.Head
	hasCycle := false

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	slow = l.Head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func (l1 *LinkedList) Zip(l2 *LinkedList) {
	if l2.Head == nil {
		return
	}
	if l1.Head == nil {
		l1.Head, l1.Tail, l1.Length = l2.Head, l2.Tail, l2.Length
		l2.Head, l2.Tail, l2.Length = nil, nil, 0
		return
	}

	dummy := &Node{}
	current := dummy

	h1 := l1.Head
	h2 := l2.Head

	for h1 != nil && h2 != nil {
		t1 := h1.Next
		t2 := h2.Next

		current.Next = h1
		h1.Prev = current

		h1.Next = h2
		h2.Prev = h1

		current = h2

		h1 = t1
		h2 = t2
	}

	if h1 != nil {
		current.Next = h1
		h1.Prev = current
		l2.Tail = l1.Tail
	} else if h2 != nil {
		current.Next = h2
		h2.Prev = current
		l1.Tail = l2.Tail
	} else {
		l1.Tail = current
	}

	l1.Head = dummy.Next
	if l1.Head != nil {
		l1.Head.Prev = nil
	}

	l1.Length += l2.Length

	l2.Head = nil
	l2.Tail = nil
	l2.Length = 0
}
