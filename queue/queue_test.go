package queue

import "testing"

func TestCreateNewQueue(t *testing.T) {
	q := NewQueue()

	if !q.isEmpty() {
		t.Error("New queue must be empty")
	}
	if q.first != nil {
		t.Error("New queue can't have node associeted in than")
	}
}

func TestPushOnQueue(t *testing.T) {
	q := NewQueue()

	testCases := []struct {
		value         int
		expectedOnTop int
	}{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 1},
	}

	for _, testCase := range testCases {
		q.push(testCase.value)

		if testCase.expectedOnTop != q.first.value {
			t.Errorf("Expected %d, but got %d", testCase.expectedOnTop, q.first.value)
		}
	}
}

func TestPopOnQueue(t *testing.T) {
	q := NewQueue()

	q.push(1)
	q.push(2)
	q.push(3)
	q.push(4)

	testCases := []struct {
		value       int
		expectedErr bool
	}{
		{1, false},
		{2, false},
		{3, false},
		{4, false},
		{0, true},
	}

	for _, testCase := range testCases {
		value, err := q.pop()

		if testCase.expectedErr && err == nil {
			t.Error("Expected error")
		} else {
			if testCase.value != value {
				t.Errorf("Expected %d,but god %d", testCase.value, value)
			}
		}
	}
}
