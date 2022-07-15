package linked_list

import "fmt"

func removeNodeIteration(input *LinkedList[string], target string) *LinkedList[string] {
	fmt.Println("Processing ", input)
	if input.Val == target {
		return input.Next
	}
	prev := input
	head := prev
	input = input.Next
	for input != nil {
		if target != input.Val {
			prev = prev.Next
		} else {
			prev.Next = input.Next
			return head
		}
		input = input.Next
	}
	return nil
}

func removeNodeRecursion(input *LinkedList[string], target string) *LinkedList[string] {
	if input.Val == target {
		return input.Next
	}
	input.Next = removeNodeRecursion(input.Next, target)
	return input
}

func TestRemove00() {
	test := NewLinkedList([]string{"a", "b", "c", "d", "e", "f"})
	test2 := NewLinkedList([]string{"a", "b", "c", "d", "e", "f"})
	target := "c"

	fmt.Println(removeNodeIteration(test, target), ". First ", target, " is removed")
	fmt.Println(removeNodeRecursion(test2, target), ". First ", target, " is removed with recursion")
}

func TestRemove01() {
	test := NewLinkedList([]string{"x", "y", "z"})
	test2 := NewLinkedList([]string{"x", "y", "z"})

	target := "z"

	fmt.Println(removeNodeIteration(test, target), ". First ", target, " is removed")
	fmt.Println(removeNodeRecursion(test2, target), ". First ", target, " is removed with recursion")
}

func TestRemove02() {
	test := NewLinkedList([]string{"q", "r", "s"})
	test2 := NewLinkedList([]string{"q", "r", "s"})
	target := "q"

	fmt.Println(removeNodeIteration(test, target), ". First ", target, " is removed")
	fmt.Println(removeNodeRecursion(test2, target), ". First ", target, " is removed with recursion")
}

func TestRemove03() {
	test := NewLinkedList([]string{"h", "i", "j", "i"})
	test2 := NewLinkedList([]string{"h", "i", "j", "i"})
	target := "i"

	fmt.Println(removeNodeIteration(test, target), ". First ", target, " is removed")
	fmt.Println(removeNodeRecursion(test2, target), ". First ", target, " is removed with recursion")
}
