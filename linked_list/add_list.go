package linked_list

import "fmt"

func addWithIteration(node1 *LinkedList[int], node2 *LinkedList[int]) *LinkedList[int] {
	result := &LinkedList[int]{}
	tail := result
	carry := 0
	for node1 != nil || node2 != nil || carry == 1 {
		val1 := 0
		if node1 != nil {
			val1 = node1.Val
			node1 = node1.Next
		}

		val2 := 0
		if node2 != nil {
			val2 = node2.Val
			node2 = node2.Next
		}

		sum := val1 + val2 + carry
		if sum > 9 {
			carry = 1
			sum = sum - 10
		} else {
			carry = 0
		}

		tail.Next = &LinkedList[int]{Val: sum}
		tail = tail.Next
	}

	return result.Next
}

func addWithRecursion(node1 *LinkedList[int], node2 *LinkedList[int], carry int) *LinkedList[int] {
	if node1 == nil && node2 == nil && carry == 0 {
		return nil
	}

	val1 := 0
	if node1 != nil {
		val1 = node1.Val
		node1 = node1.Next
	}
	val2 := 0
	if node2 != nil {
		val2 = node2.Val
		node2 = node2.Next
	}

	sum := val1 + val2 + carry
	if sum > 9 {
		carry = 1
		sum = sum - 10
	} else {
		carry = 0
	}

	result := &LinkedList[int]{Val: sum}
	result.Next = addWithRecursion(node1, node2, carry)
	return result
}

func TestAdd00() {
	testA := NewLinkedList([]int{1, 2, 6})
	testB := NewLinkedList([]int{4, 5, 3})

	fmt.Println(addWithIteration(testA, testB))
	fmt.Println(addWithRecursion(testA, testB, 0))
}

func TestAdd01() {
	testA := NewLinkedList([]int{1, 4, 5, 7})
	testB := NewLinkedList([]int{2, 3})

	fmt.Println(addWithIteration(testA, testB))
	fmt.Println(addWithRecursion(testA, testB, 0))
}

func TestAdd02() {
	testA := NewLinkedList([]int{9, 3})
	testB := NewLinkedList([]int{7, 4})

	fmt.Println(addWithIteration(testA, testB))
	fmt.Println(addWithRecursion(testA, testB, 0))
}

func TestAdd03() {
	testA := NewLinkedList([]int{9, 8})
	testB := NewLinkedList([]int{7, 4})

	fmt.Println(addWithIteration(testA, testB))
	fmt.Println(addWithRecursion(testA, testB, 0))
}

func TestAdd04() {
	testA := NewLinkedList([]int{9, 9, 9})
	testB := NewLinkedList([]int{6})

	fmt.Println(addWithIteration(testA, testB))
	fmt.Println(addWithRecursion(testA, testB, 0))
}
