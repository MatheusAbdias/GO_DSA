package queue

import "fmt"

type Queue struct {
	len   int
	first *Node
}

type Node struct {
	value int
	next  *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

func (q *Queue) isEmpty() bool {
	return q.len == 0
}

func (q *Queue) push(value int) {
	newNode := NewNode(value)
	if q.isEmpty() {
		q.first = newNode
	} else {
		node := q.first
		for node.next != nil {
			node = node.next
		}

		node.next = newNode

	}
	q.len++
}

func (q *Queue) pop() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("Cannot pop in empty queue")
	} else {
		value := q.first.value
		q.first = q.first.next
		q.len--
		return value, nil
	}
}
