package array

import "testing"

func TestCreateNewArray(t *testing.T) {

	array := NewArray()

	if array.length != 0 {
		t.Error("Expected inicialize a empty array")
	}

}

func TestAppendInArray(t *testing.T) {

	value := 1
	array := NewArray()

	array.Append(value)

	if array.length != value {
		t.Error("Expected append a new value in array")
	}
	if array.data[0] != value {
		t.Error("Expected append a new value in array")
	}
}

func TestGetInArray(t *testing.T) {

	value := 1
	array := NewArray()

	array.Append(value)

	result, err := array.Get(0)

	if err != nil {
		t.Error("Expected get a value in array")
	}
	if result != value {
		t.Error("Expected get a value in array")
	}
}

func TestGetInArrayWithIndexOutOfRange(t *testing.T) {

	array := NewArray()

	_, err := array.Get(0)

	if err == nil {
		t.Error("Expected get a value in array")
	}
}

func TestPopInArray(t *testing.T) {

	value := 1
	array := NewArray()

	array.Append(value)

	result, err := array.pop()

	if err != nil {
		t.Error("Expected pop a value in array")
	}
	if result != value {
		t.Error("Expected pop a value in array")
	}
	if array.length != 0 {
		t.Error("Expected pop a value in array")
	}
}

func TestPopInArrayWithEmptyArray(t *testing.T) {

	array := NewArray()

	_, err := array.pop()

	if err == nil {
		t.Error("Expected pop a value in array")
	}
}
