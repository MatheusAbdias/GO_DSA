package array

func compareArray(a *Array, expected []int) bool {
	if a.len() != len(expected) {
		return false
	}
	for i, val := range expected {
		if a.data[i] != val {
			return false
		}
	}
	return true
}
