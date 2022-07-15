package linked_list

import "fmt"

func PrintLinkedListWithIteration(current *LinkedList[int]) {
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func PrintLinkedListRecursion(current *LinkedList[int]) {
	if current == nil {
		return
	}
	fmt.Println(current.Val)
	PrintLinkedListRecursion(current.Next)
}
