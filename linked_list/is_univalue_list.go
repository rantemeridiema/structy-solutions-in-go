package linked_list

import "fmt"

func IsUnivalueIteration[T comparable](input *LinkedList[T]) bool {
	value := input.Val
	for input != nil {
		if input.Val != value {
			return false
		}
		input = input.Next
	}
	return true
}

func IsUnivalueRecursion[T comparable](input *LinkedList[T], value T) bool {
	if input == nil {
		return true
	}
	if input.Val != value {
		return false
	}
	return IsUnivalueRecursion(input.Next, value)
}

func TestUnivalue() {
	test := NewLinkedList([]int{7, 7, 7})
	test2 := NewLinkedList([]int{7, 7, 4})
	test3 := NewLinkedList([]int{2, 2, 2, 2, 2})
	test4 := NewLinkedList([]int{2, 2, 3, 3, 2})
	test5 := NewLinkedList([]string{"z"})

	fmt.Println(IsUnivalueIteration(test))
	fmt.Println(IsUnivalueRecursion(test, test.Val))

	fmt.Println(IsUnivalueIteration(test2))
	fmt.Println(IsUnivalueRecursion(test2, test2.Val))

	fmt.Println(IsUnivalueIteration(test3))
	fmt.Println(IsUnivalueRecursion(test3, test3.Val))

	fmt.Println(IsUnivalueIteration(test4))
	fmt.Println(IsUnivalueRecursion(test4, test4.Val))

	fmt.Println(IsUnivalueIteration(test5))
	fmt.Println(IsUnivalueRecursion(test5, test5.Val))
}
