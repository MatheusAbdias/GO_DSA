package linkedlist

import "fmt"

type LinkedList struct {
	length int
	next   *Node
}

type Node struct {
	value int
	next  *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Length() int {
	return l.length
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func (l *LinkedList) Append(value int) {
	newNode := NewNode(value)
	if l.length == 0 {
		l.next = newNode
	} else {
		lastNode := l.next
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = newNode
	}

	l.length++
}

func (l *LinkedList) Pop() (int, error) {

	if l.length == 0 {
		return 0, fmt.Errorf("Cannot pop in empty linked list")
	}
	if l.length == 1 {
		value := l.next.value
		l.next = nil
		l.length--
		return value, nil
	}
	node := l.next
	for node.next.next != nil {
		node = node.next
	}

	value := node.next.value
	node.next = nil
	l.length--
	return value, nil
}

func (l *LinkedList) popIndex(index int) (int, error) {
	if index >= l.length || l.length <= 0 {
		return 0, fmt.Errorf("Invalid index, out off range")
	}

	if index == 0 {
		value := l.next.value
		if l.next.next != nil {
			l.next = l.next.next
		} else {
			l.next = nil
		}
		l.length--
		return value, nil
	}

	node := l.next
	for i := 0; i < index-1; i++ {
		node = node.next
	}
	value := node.next.value

	if node.next.next != nil {
		node.next = node.next.next
	} else {
		node.next = nil
	}
	l.length--
	return value, nil
}
