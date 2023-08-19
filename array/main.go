package array

import "fmt"

type Array struct {
	data   []int
	length int
}

func NewArray() *Array {
	return &Array{}
}

func (a *Array) Append(value int) {
	a.data = append(a.data, value)
	a.length++
}

func (a *Array) Get(index int) (int, error) {
	if index < 0 || index >= a.length {
		return 0, fmt.Errorf("Index out of range")
	}

	return a.data[index], nil
}

func (a *Array) pop() (int, error) {
	if a.length <= 0 {
		return 0, fmt.Errorf("Empty array cannot use pop")
	}
	last := a.data[a.length-1]
	a.length--

	return last, nil
}

func (a *Array) Length() int {
	return a.length
}
