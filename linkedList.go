package main

import "errors"

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
