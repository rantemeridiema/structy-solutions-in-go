package linked_list

import "fmt"

func reverseWithIteration[T any](input *LinkedList[T]) *LinkedList[T] {
	current := input
	var prev *LinkedList[T]
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func reverseWithRecursion[T any](input *LinkedList[T], prev *LinkedList[T]) *LinkedList[T] {
	if input == nil {
		return prev
	}
	next := input.Next
	input.Next = prev
	return reverseWithRecursion(next, input)
}

func TestReverse() {
	f := LinkedList[string]{Val: "f"}
	e := LinkedList[string]{Val: "e", Next: &f}
	d := LinkedList[string]{Val: "d", Next: &e}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(reverseWithIteration(&a))
	fmt.Println(reverseWithRecursion(&f, nil)) //reverse again
}

func TestReverseOneNode() {
	f := LinkedList[string]{Val: "f"}

	fmt.Println(reverseWithIteration(&f))
	fmt.Println(reverseWithRecursion(&f, nil))
}
