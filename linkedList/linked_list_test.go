package linkedlist

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	l := NewLinkedList()

	if l.length > 0 {
		t.Errorf("Expected new linked list is empty.")
	}

	if l.next != nil {
		t.Error("Expected next element in new linke list is Null")
	}
}

func TestAppendInLinkedList(t *testing.T) {
	l := NewLinkedList()
	l.Append(10)

	if l.length == 0 {
		t.Error("Expected length grater than zero when append element")
	}
	if l.next == nil {
		t.Error("Expected next node is not nil")
	}
}

func TestPopLinkedList(t *testing.T) {
	l := NewLinkedList()
	value := 10
	l.Append(value)

	popedValue, err := l.Pop()

	if err != nil {
		t.Error("Not expected error when pop value in not blank linked list")
	}
	if popedValue != value {
		t.Error("Exeped popedValue equals a appended value")
	}

}

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()

	if list.Length() != 0 {
		t.Errorf("Expected length 0, but got %d", list.Length())
	}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	if list.Length() != 3 {
		t.Errorf("Expected length 3, but got %d", list.Length())
	}

	value, err := list.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected value 30, but got %d", value)
	}

	if list.Length() != 2 {
		t.Errorf("Expected length 2, but got %d", list.Length())
	}

	value, err = list.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 20 {
		t.Errorf("Expected value 20, but got %d", value)
	}

	if list.Length() != 1 {
		t.Errorf("Expected length 1, but got %d", list.Length())
	}

	value, err = list.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected value 10, but got %d", value)
	}

	if list.Length() != 0 {
		t.Errorf("Expected length 0, but got %d", list.Length())
	}

	_, err = list.Pop()
	if err == nil {
		t.Errorf("Expected error for popping from an empty list, but got no error")
	}
}
func TestLinkedListPopIndex(t *testing.T) {
	list := NewLinkedList()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)

	tests := []struct {
		index        int
		expectedList []int
		expectedErr  bool
	}{
		{2, []int{10, 20, 40, 50}, false},
		{0, []int{20, 40, 50}, false},
		{1, []int{20, 50}, false},
		{1, []int{20}, false},
		{0, []int{}, false},
		{0, []int{}, true},
	}

	for _, test := range tests {
		_, err := list.popIndex(test.index)

		if (err != nil) != test.expectedErr {
			t.Errorf("Error mismatch. Expected: %v, Got: %v", test.expectedErr, err != nil)
		}

		if !test.expectedErr && !compareList(list, test.expectedList) {
			t.Errorf("List mismatch for index %d. Expected: %v, Got: %v", test.index, test.expectedList, listToArray(list))
		}
	}
}

func compareList(list *LinkedList, expected []int) bool {
	if list.Length() != len(expected) {
		return false
	}

	node := list.next
	for _, val := range expected {
		if node.value != val {
			return false
		}
		node = node.next
	}
	return true
}

func listToArray(list *LinkedList) []int {
	arr := []int{}
	node := list.next
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}
	return arr
}
