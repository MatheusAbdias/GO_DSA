package bst

import (
	"testing"
)

func TestCreateBST(t *testing.T) {
	bst := NewBST()

	if bst.head != nil || !bst.IsEmpty() {
		t.Error("Expected empty head when start new BST")
	}
}

func TestNavInPreOrder(t *testing.T) {
	/*       [1]
		 .[-1]     [3]
	  [-2] [0] [2] [4]
	*/

	b := NewBST()
	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	expected := []int{1, -1, -2, 0, 3, 2, 4}
	result := []int{}
	NavPreOrder(b.head, &result)

	for index, value := range result {
		if value != expected[index] {
			t.Errorf("Expected value %d but receive %d", expected[index], value)
		}
	}
}

func TestNavInPostOrder(t *testing.T) {
	/*       [1]
		 .[-1]     [3]
	  [-2] [0] [2] [4]
	*/

	b := NewBST()
	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	expected := []int{-2, 0, -1, 2, 4, 3, 1}
	result := []int{}

	NavInPostOrder(b.head, &result)

	for index, value := range result {
		if value != expected[index] {
			t.Errorf("Expected value %d but receive %d", expected[index], value)
		}
	}
}

func TestSearchInBST(t *testing.T) {
	b := NewBST()

	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	testCases := []struct {
		Value int
		Find  bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{5, false},
	}

	for _, testCase := range testCases {
		node, err := b.head.Search(testCase.Value)

		if testCase.Find && node.value != testCase.Value {
			t.Errorf("Expected value %d, but receive %d", testCase.Value, node.value)
		} else {
			if !testCase.Find && err == nil {
				t.Error("Is expected return a error when value is not found")
			}
		}
	}
}

func TestMinInBst(t *testing.T) {
	b := NewBST()

	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	min := b.head.Min()
	minValue := -2
	if min.value != minValue {
		t.Errorf("Expected %d, but receive %d", minValue, min.value)
	}
}

func TestMaxInBst(t *testing.T) {
	b := NewBST()

	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	max := b.head.Max()
	maxValue := 4
	if max.value != maxValue {
		t.Errorf("Expected %d, but receive %d", maxValue, max.value)
	}
}

func TestHeightInBst(t *testing.T) {
	/*       [1]
		 .[-1]     [3]
	  [-2] [0] [2] [4]
	*/

	b := NewBST()

	b.Insert(1)
	b.Insert(3)
	b.Insert(2)
	b.Insert(4)
	b.Insert(-1)
	b.Insert(-2)
	b.Insert(0)

	height := b.head.Depth()
	expectedHeight := 3

	if expectedHeight != height {
		t.Errorf("Expected %d, but receive %d", expectedHeight, height)
	}
}

func TestDeleteInBST(t *testing.T) {
	b := NewBST()

	b.Insert(5)
	b.Insert(3)
	b.Insert(7)
	b.Insert(2)
	b.Insert(4)
	b.Insert(6)
	b.Insert(8)

	b.head.Delete(2)
	if _, err := b.head.Search(2); err == nil {
		t.Error("Expected error for deleted value 2")
	}

	b.head.Delete(7)
	if _, err := b.head.Search(7); err == nil {
		t.Error("Expected error for deleted value 7")
	}

	b.head.Delete(5)
	if _, err := b.head.Search(5); err == nil {
		t.Error("Expected error for deleted value 5")
	}

	testCases := []struct {
		Value int
		Find  bool
	}{
		{3, true},
		{4, true},
		{6, true},
		{8, true},
	}

	for _, testCase := range testCases {
		_, err := b.head.Search(testCase.Value)
		if testCase.Find && err != nil {
			t.Errorf("Expected value %d, but not found", testCase.Value)
		} else if !testCase.Find && err == nil {
			t.Errorf("Value %d should not be present, but found", testCase.Value)
		}
	}
}
