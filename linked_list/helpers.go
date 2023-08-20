package linkedlist

func compareList(list *LinkedList, expected []int) bool {
	if list.len() != len(expected) {
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
