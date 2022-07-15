package linked_list

import "fmt"

func getNodeValueWithIteration(input *LinkedList[string], targetIndex int) string {
	var result string
	current := input
	index := 0
	for current != nil && index < targetIndex {
		current = current.Next
		index++
	}
	if index == targetIndex {
		result = current.Val
	}
	return result
}

func getNodeWithRecursion(input *LinkedList[string], targetIndex int) string {
	if input == nil {
		return ""
	}
	if targetIndex == 0 {
		return input.Val
	}
	return getNodeWithRecursion(input.Next, targetIndex-1)
}

func TestGetNodeValue00() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(getNodeValueWithIteration(&a, 2))
	fmt.Println(getNodeWithRecursion(&a, 2))
}

func TestGetNodeValue01() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(getNodeValueWithIteration(&a, 3))
	fmt.Println(getNodeWithRecursion(&a, 3))
}

func TestGetNodeValue02() {
	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	fmt.Println(getNodeValueWithIteration(&a, 7))
	fmt.Println(getNodeWithRecursion(&a, 7))
}

func TestGetNodeValue03() {
	node2 := LinkedList[string]{Val: "mango"}
	node1 := LinkedList[string]{Val: "banana", Next: &node2}

	fmt.Println(getNodeValueWithIteration(&node1, 0))
	fmt.Println(getNodeWithRecursion(&node1, 0))
}

func TestGetNodeValue04() {
	node2 := LinkedList[string]{Val: "mango"}
	node1 := LinkedList[string]{Val: "banana", Next: &node2}

	fmt.Println(getNodeValueWithIteration(&node1, 1))
	fmt.Println(getNodeWithRecursion(&node1, 1))
}
