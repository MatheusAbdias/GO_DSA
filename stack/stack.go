package stack

import "fmt"

type Stack struct {
	len int
	top *Node
}

type Node struct {
	value   int
	previus *Node
}

func NewStack() *Stack { return &Stack{} }

func NewNode(value int, previus *Node) *Node { return &Node{value, previus} }

func (s *Stack) isEmpty() bool { return s.len == 0 }

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("Cannot pop in empty stack")
	}

	value := s.top.value
	s.top = s.top.previus
	s.len--
	return value, nil
}

func (s *Stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("Cannot peak in empty stack")
	}
	return s.top.value, nil
}

func (s *Stack) clear() {
	s.top = nil
	s.len = 0
}

func (s *Stack) push(value int) {
	s.top = NewNode(value, s.top)
	s.len++
}
