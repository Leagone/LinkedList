package main

import (
	"errors"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
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
	}

	l.Length++

}

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{
		Value: value,
	}
	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) Pop() (int, error) {

	if l.Head == nil {
		return 0, errors.New("Cannot pop from empty list")
	}

	var val int

	if l.Head.Next == nil {
		val = l.Head.Value
		l.Head = nil
	} else {
		current := l.Head
		for current.Next.Next != nil {
			current = current.Next
		}
		val = current.Next.Value
		current.Next = nil
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

	if index < 0 || index > l.Length {
		return errors.New("Out of Bounds")
	}

	newNode := &Node{
		Value: value,
	}

	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		l.Length++
		return nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.Length++
	return nil

}

func (l *LinkedList) RemoveAt(index int) error {
	if index < 0 || index >= l.Length {
		return errors.New("Out of Bounds")
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	current := l.Head

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
	l.Length--
	return nil
}

func (l *LinkedList) RemoveByValue(value int) error {

	if l.Head == nil {
		return errors.New("Cannot remove from empty list")
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		return nil
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.Length--
			return nil
		}
		current = current.Next
	}
	return errors.New("Value not Found")
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
		if condition(current.Value) {
			tail.Next = current
			tail = tail.Next
			newLenght++
		}
		current = current.Next
	}

	tail.Next = nil

	l.Head = dummy.Next
	l.Length = newLenght
}

func (l *LinkedList) Reverse() {

	var previous *Node
	current := l.Head

	for current != nil {
		tail := current.Next
		current.Next = previous
		previous = current
		current = tail
	}

	l.Head = previous
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

	if n <= 0 {
		return errors.New("Out of Bounds")
	}

	start := &Node{
		Next: l.Head,
	}

	slow := start
	fast := start

	for i := 0; i <= n; i++ {
		if fast == nil {
			return errors.New("Out of Bounds")
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	l.Length--
	l.Head = start.Next

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

	dummy := &Node{}

	current := dummy

	h1 := l1.Head
	h2 := l2.Head

	for h1 != nil && h2 != nil {
		t1 := h1.Next
		t2 := h2.Next

		h1.Next = nil
		h2.Next = nil

		current.Next = h1
		current.Next.Next = h2

		current = h2

		h1 = t1
		h2 = t2

	}

	if h1 != nil {
		current.Next = h1
	}

	if h2 != nil {
		current.Next = h2
	}

	l1.Head = dummy.Next
	l1.Length = l1.Length + l2.Length

}
