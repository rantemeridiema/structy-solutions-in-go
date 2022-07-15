package linked_list

func LinkedListToArrayIteration(current *LinkedList[int]) []int {
	var result []int
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func LinkedListToArrayRecursion(current *LinkedList[int], result []int) []int {
	if current == nil {
		return result
	}
	result = append(result, current.Val)
	return LinkedListToArrayRecursion(current.Next, result)
}
