package stack

import "testing"

func TestPushInStack(t *testing.T) {
	s := NewStack()

	testCases := []struct {
		value         int
		expectedOnTop int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}

	for _, testCase := range testCases {

		s.push(testCase.value)
		if s.top.value != testCase.expectedOnTop {
			t.Errorf("Expected on top %d,but got %d", testCase.expectedOnTop, s.top.value)
		}
	}
}

func TestPopInStack(t *testing.T) {
	s := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	testCases := []struct {
		value         int
		expectedOnTop int
		expectedErr   bool
	}{
		{4, 3, false},
		{3, 2, false},
		{2, 1, false},
		{1, 0, false},
		{0, 0, false},
	}

	for _, testCase := range testCases {

		value, err := s.pop()

		if testCase.expectedErr && err == nil {
			t.Error("Exepected err")
		} else {
			if testCase.value != value {
				t.Errorf("Expected value on pop %d, but got %d.", testCase.value, value)
			}
			if testCase.expectedOnTop != 0 && testCase.expectedOnTop != s.top.value {
				t.Errorf("Expected value on top %d, but got %d.", testCase.expectedOnTop, s.top.value)
			}
		}
	}
}

func TestClearStack(t *testing.T) {
	s := NewStack()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)

	s.clear()

	if !s.isEmpty() {
		t.Error("Expected empty stack after clear")
	}
	if s.top != nil {
		t.Error("Expected nothing on top of the list")
	}

}

func TestPeekInStack(t *testing.T) {
	s := NewStack()

	s.push(1)
	s.push(2)

	testCases := []struct {
		value       int
		expectedErr bool
	}{
		{2, false},
		{1, false},
		{0, true},
	}

	for _, testCase := range testCases {
		value, err := s.peek()
		s.pop()

		if testCase.expectedErr && err == nil {
			t.Error("Expected erro")
		} else {
			if testCase.value != value {
				t.Errorf("Expected value %d,but got %d", testCase.value, value)
			}
		}
	}
}
