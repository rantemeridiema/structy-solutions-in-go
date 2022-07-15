package linked_list

import "fmt"

func sumListIteration(head *LinkedList[int]) int {
	var result int
	var current = head
	for current != nil {
		result = result + current.Val
		current = current.Next
	}
	return result
}

func sumListRecursion(head *LinkedList[int]) int {
	if head == nil {
		return 0
	}
	return head.Val + sumListIteration(head.Next)
}

func TestSumList00() {
	var e = LinkedList[int]{Val: 7}
	var d = LinkedList[int]{-1, &e}
	var c = LinkedList[int]{3, &d}
	var b = LinkedList[int]{8, &c}
	var a = LinkedList[int]{2, &b}

	fmt.Println(sumListIteration(&a))
	fmt.Println(sumListRecursion(&a))
}

func TestSumList01() {
	var b = LinkedList[int]{Val: 4}
	var a = LinkedList[int]{38, &b}

	fmt.Println(sumListIteration(&a))
	fmt.Println(sumListRecursion(&a))
}

func TestSumList02() {
	var a = LinkedList[int]{Val: 100}

	fmt.Println(sumListIteration(&a))
	fmt.Println(sumListRecursion(&a))
}

func TestSumList03() {

	fmt.Println(sumListIteration(nil))
	fmt.Println(sumListRecursion(nil))
}
