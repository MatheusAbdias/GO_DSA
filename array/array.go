package array

import "fmt"

type Array struct {
	data   []int
	length int
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) append(value int) {
	a.data = append(a.data, value)
	a.length++
}

func (a *Array) isEmpty() bool {
	return a.length == 0
}
func (a *Array) Get(index int) (int, error) {
	if index < 0 || index >= a.length {
		return 0, fmt.Errorf("Index out of range")
	}

	return a.data[index], nil
}

func (a *Array) pop() (int, error) {
	if a.isEmpty() {
		return 0, fmt.Errorf("Empty array cannot use pop")
	}
	last := a.data[a.length-1]
	a.length--
	a.data = a.data[:a.length]

	return last, nil
}

func (a *Array) popIndex(index int) (int, error) {
	if index >= a.length {
		return 0, fmt.Errorf("Indext cant be greater than array length")
	}

	value := a.data[index]
	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	a.data = a.data[:a.length]
	return value, nil
}

func (a *Array) len() int {
	return a.length
}
