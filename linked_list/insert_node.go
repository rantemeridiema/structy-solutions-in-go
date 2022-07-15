package linked_list

import "fmt"

func insertNodeRecursion(input *LinkedList[string], targetVal string, targetIndex int) *LinkedList[string] {
	if targetIndex == 0 {
		targetLinkedList := &LinkedList[string]{targetVal, input}
		return targetLinkedList
	}
	input.Next = insertNodeRecursion(input.Next, targetVal, targetIndex-1)
	return input
}

func insertNodeIteration(input *LinkedList[string], targetVal string, targetIndex int) *LinkedList[string] {
	if targetIndex == 0 {
		return &LinkedList[string]{targetVal, input}
	}
	current := input
	for {
		if targetIndex == 1 {
			next := current.Next
			current.Next = &LinkedList[string]{targetVal, next}
			return input
		}
		current = current.Next
		targetIndex--
	}
}

func TestInsert00() {
	input := NewLinkedList([]string{"a", "b", "c", "d"})
	input2 := NewLinkedList([]string{"a", "b", "c", "d"})

	fmt.Println(insertNodeRecursion(input, "x", 2))
	fmt.Println(insertNodeIteration(input2, "x", 2))
}

func TestInsert01() {
	input := NewLinkedList([]string{"a", "b", "c", "d"})
	input2 := NewLinkedList([]string{"a", "b", "c", "d"})

	fmt.Println(insertNodeRecursion(input, "v", 3))
	fmt.Println(insertNodeIteration(input2, "v", 3))
}

func TestInsert02() {
	input := NewLinkedList([]string{"a", "b", "c", "d"})
	input2 := NewLinkedList([]string{"a", "b", "c", "d"})

	fmt.Println(insertNodeRecursion(input, "m", 4))
	fmt.Println(insertNodeIteration(input2, "m", 4))
}

func TestInsert03() {
	input := NewLinkedList([]string{"a", "b"})
	input2 := NewLinkedList([]string{"a", "b"})

	fmt.Println(insertNodeRecursion(input, "z", 0))
	fmt.Println(insertNodeRecursion(input2, "z", 0))
}
