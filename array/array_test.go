package array

import (
	"testing"
)

func TestCreateNewArray(t *testing.T) {

	array := NewArray()

	if array.length != 0 {
		t.Error("Expected inicialize a empty array")
	}

}

func TestAppendInArray(t *testing.T) {
	a := NewArray()
	testCases := []struct {
		value        int
		expectedList []int
		expectedErr  bool
	}{
		{1, []int{1}, false},
		{2, []int{1, 2}, false},
		{3, []int{1, 2, 3}, false},
	}

	for _, testCase := range testCases {
		a.append(testCase.value)

		if testCase.expectedErr == compareArray(a, testCase.expectedList) {
			t.Errorf("Expected array %v, but got %v", testCase.expectedList, a.data)
		}
	}
}
func TestGetInArray(t *testing.T) {
	a := NewArray()
	a.append(1)
	a.append(2)
	a.append(3)
	a.append(4)

	testCases := []struct {
		index       int
		value       int
		expectedErr bool
	}{
		{0, 1, false},
		{1, 2, false},
		{2, 3, false},
		{3, 4, false},
		{4, 0, true},
	}

	for _, testCase := range testCases {
		value, err := a.Get(testCase.index)

		if testCase.expectedErr && err == nil {
			t.Error("Exepected Error")
		} else {
			if value != testCase.value {
				t.Errorf("Expected value %d, but got %d", testCase.value, value)
			}
		}
	}
}

func TestPopInArray(t *testing.T) {
	a := NewArray()
	a.append(1)
	a.append(2)
	a.append(3)
	a.append(4)
	testCases := []struct {
		value       int
		expectedErr bool
	}{
		{4, false},
		{3, false},
		{2, false},
		{1, false},
		{0, true},
	}

	for _, testCase := range testCases {
		value, err := a.pop()

		if testCase.expectedErr && err == nil {
			t.Error("Expected error")
		} else {
			if value != testCase.value {
				t.Errorf("Expected value %d, but got %d", testCase.value, value)
			}
		}
	}
}

func TestPopIndexInArray(t *testing.T) {
	a := NewArray()
	a.append(1)
	a.append(2)
	a.append(3)
	a.append(4)
	testCases := []struct {
		index        int
		expectedList []int
		expectedErr  bool
	}{
		{0, []int{2, 3, 4}, false},
		{1, []int{2, 4}, false},
		{1, []int{2}, false},
		{0, []int{}, false},
		{0, nil, true},
	}

	for _, testCase := range testCases {
		_, err := a.popIndex(testCase.index)

		if testCase.expectedErr && err == nil {
			t.Error("Expected error")
		} else {
			if !compareArray(a, testCase.expectedList) {
				t.Errorf("Expected value %v, but got %v", testCase.expectedList, a.data)
			}
		}
	}
}
