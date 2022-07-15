package linked_list

import "fmt"

func zipWithIteration[T any](node1 *LinkedList[T], node2 *LinkedList[T]) *LinkedList[T] {
	result := node1
	for {
		next1 := node1.Next
		next2 := node2.Next
		node1.Next = node2
		if next1 == nil {
			break
		}
		node2.Next = next1
		if next2 == nil {
			break
		}
		node2 = next2
		node1 = next1
	}
	return result
}

func zipWithIRecursion[T any](node1 *LinkedList[T], node2 *LinkedList[T]) *LinkedList[T] {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	next1 := node1.Next
	next2 := node2.Next
	node1.Next = node2
	node2.Next = zipWithIRecursion(next1, next2)
	return node1
}

func TestZip00Iteration() {

	c := LinkedList[string]{Val: "c"}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	three := LinkedList[string]{Val: "3"}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIteration(&a, &one))
}

func TestZip00Recursion() {

	c := LinkedList[string]{Val: "c"}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	three := LinkedList[string]{Val: "3"}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIRecursion(&a, &one))
}

func TestZip01Iteration() {

	c := LinkedList[string]{Val: "c"}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	four := LinkedList[string]{Val: "4"}
	three := LinkedList[string]{Val: "3", Next: &four}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIteration(&a, &one))
}

func TestZip01Recursion() {

	c := LinkedList[string]{Val: "c"}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	four := LinkedList[string]{Val: "4"}
	three := LinkedList[string]{Val: "3", Next: &four}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIRecursion(&a, &one))
}

func TestZip02Iteration() {

	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	three := LinkedList[string]{Val: "3"}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIteration(&a, &one))
}

func TestZip02Recursion() {

	d := LinkedList[string]{Val: "d"}
	c := LinkedList[string]{Val: "c", Next: &d}
	b := LinkedList[string]{Val: "b", Next: &c}
	a := LinkedList[string]{Val: "a", Next: &b}

	three := LinkedList[string]{Val: "3"}
	two := LinkedList[string]{Val: "2", Next: &three}
	one := LinkedList[string]{Val: "1", Next: &two}

	fmt.Println(zipWithIRecursion(&a, &one))
}
