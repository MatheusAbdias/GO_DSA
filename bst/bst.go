package bst

import (
	"errors"
)

type BST struct {
	head   *NoBST
	length int
}

type NoBST struct {
	value int
	left  *NoBST
	right *NoBST
}

func NewBST() *BST {
	return &BST{}
}

func NewNoBST(value int) *NoBST {
	return &NoBST{value: value}
}

func (b *BST) IsEmpty() bool {
	return b.length == 0
}

func (n *NoBST) Insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = NewNoBST(value)
		} else {
			n.left.Insert(value)
		}
	} else if value > n.value {
		if n.right == nil {
			n.right = NewNoBST(value)
		} else {
			n.right.Insert(value)
		}
	}
}

func (b *BST) Insert(value int) {
	if b.head == nil {
		b.head = NewNoBST(value)
	} else {
		b.head.Insert(value)
	}

	b.length++
}

func (n *NoBST) Search(value int) (*NoBST, error) {
	if n == nil {
		return nil, errors.New("Value not found")
	}

	if n.value == value {
		return n, nil
	}

	if value < n.value {
		return n.left.Search(value)
	} else {
		return n.right.Search(value)
	}
}

func (n *NoBST) Delete(value int) *NoBST {
	if n == nil {
		return n
	}

	if value < n.value {
		n.left = n.left.Delete(value)
	} else if value > n.value {
		n.right = n.right.Delete(value)
	} else {

		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		successor := n.right.Min()
		n.value = successor.value
		n.right = n.right.Delete(successor.value)
	}

	return n
}

func NavPreOrder(n *NoBST, result *[]int) {
	if n != nil {
		*result = append(*result, n.value)
		NavPreOrder(n.left, result)
		NavPreOrder(n.right, result)
	}
}

func NavInOrder(n *NoBST, result *[]int) {
	if n != nil {
		NavInOrder(n.right, result)
		*result = append(*result, n.value)
		NavInOrder(n.left, result)
	}
}

func NavInPostOrder(n *NoBST, result *[]int) {
	if n != nil {
		NavInPostOrder(n.left, result)
		NavInPostOrder(n.right, result)
		*result = append(*result, n.value)
	}
}

func (n *NoBST) Min() *NoBST {
	node := n
	for node.left != nil {
		node = node.left
	}

	return node
}

func (n *NoBST) Max() *NoBST {
	node := n
	for node.right != nil {
		node = node.right
	}

	return node
}

func (n *NoBST) Clear() {
	n = nil
}

func (n *NoBST) Depth() int {
	if n == nil {
		return 0
	}

	leftHeight := n.left.Depth()
	rightHeight := n.right.Depth()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}

}
