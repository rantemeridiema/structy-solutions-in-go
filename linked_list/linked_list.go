package linked_list

import "fmt"

type LinkedList[T any] struct {
	Val  T
	Next *LinkedList[T]
}

func (linkedList *LinkedList[T]) String() string {
	current := linkedList
	var result string
	for current != nil {
		result = fmt.Sprint(result, current.Val)
		if current.Next != nil {
			result = fmt.Sprint(result, " -> ")
		}
		current = current.Next
	}
	return result
}

func NewLinkedList[T any](input []T) *LinkedList[T] {
	result := &LinkedList[T]{}
	tail := result
	for _, data := range input {
		tail.Next = &LinkedList[T]{Val: data}
		tail = tail.Next
	}
	return result.Next
}
