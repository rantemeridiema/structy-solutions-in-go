package linked_list

import "fmt"

func findWithIteration[T comparable](input *LinkedList[T], target T) bool {
	current := input
	for current != nil {
		if current.Val == target {
			return true
		}
		current = current.Next
	}
	return false
}

func findWithRecursion[T comparable](input *LinkedList[T], target T) bool {
	if input == nil {
		return false
	}
	if input.Val == target {
		return true
	}
	return findWithRecursion(input.Next, target)
}

func TestFindList00() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(findWithIteration(&a, "c"))
	fmt.Println(findWithRecursion(&a, "c"))
}

func TestFindList01() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(findWithIteration(&a, "d"))
	fmt.Println(findWithRecursion(&a, "d"))
}

func TestFindList02() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(findWithIteration(&a, "q"))
	fmt.Println(findWithRecursion(&a, "q"))
}

func TestFindList03() {
	node2 := LinkedList[string]{Val: "lineli"}
	node1 := LinkedList[string]{Val: "jason", Next: &node2}

	fmt.Println(findWithIteration(&node1, "jason"))
	fmt.Println(findWithRecursion(&node1, "jason"))
}

func TestFindList04() {
	node1 := LinkedList[int]{Val: 42}

	fmt.Println(findWithIteration(&node1, 42))
	fmt.Println(findWithRecursion(&node1, 42))
}

func TestFindList05() {
	node1 := LinkedList[int]{Val: 42}

	fmt.Println(findWithIteration(&node1, 100))
	fmt.Println(findWithRecursion(&node1, 100))
}
