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
	l.append(10)

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
	l.append(value)

	popedValue, err := l.pop()

	if err != nil {
		t.Error("Not expected error when pop value in not blank linked list")
	}
	if popedValue != value {
		t.Error("Exeped popedValue equals a appended value")
	}

}

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()

	if list.len() != 0 {
		t.Errorf("Expected length 0, but got %d", list.len())
	}

	list.append(10)
	list.append(20)
	list.append(30)

	if list.len() != 3 {
		t.Errorf("Expected length 3, but got %d", list.len())
	}

	value, err := list.pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected value 30, but got %d", value)
	}

	if list.len() != 2 {
		t.Errorf("Expected length 2, but got %d", list.len())
	}

	value, err = list.pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 20 {
		t.Errorf("Expected value 20, but got %d", value)
	}

	if list.len() != 1 {
		t.Errorf("Expected length 1, but got %d", list.len())
	}

	value, err = list.pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected value 10, but got %d", value)
	}

	if list.len() != 0 {
		t.Errorf("Expected length 0, but got %d", list.len())
	}

	_, err = list.pop()
	if err == nil {
		t.Errorf("Expected error for popping from an empty list, but got no error")
	}
}
func TestLinkedListPopIndex(t *testing.T) {
	list := NewLinkedList()

	list.append(10)
	list.append(20)
	list.append(30)
	list.append(40)
	list.append(50)

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

func TestLinkedListInsertIndex(t *testing.T) {
	list := NewLinkedList()

	tests := []struct {
		index        int
		value        int
		expectedList []int
		expectedErr  bool
	}{
		{0, 10, []int{10}, false},
		{0, 20, []int{20, 10}, false},
		{1, 30, []int{20, 30, 10}, false},
		{3, 40, []int{20, 30, 10, 40}, false},
		{5, 50, []int{20, 30, 10, 40}, true},
		{4, 50, []int{20, 30, 10, 40, 50}, false},
	}

	for _, test := range tests {
		err := list.insertIndex(test.index, test.value)

		if (err != nil) != test.expectedErr {
			t.Errorf("Error mismatch. Expected: %v, Got: %v", test.expectedErr, err != nil)
		}

		if !test.expectedErr && !compareList(list, test.expectedList) {
			t.Errorf("List mismatch for index %d. Expected: %v, Got: %v", test.index, test.expectedList, listToArray(list))
		}
	}
}
