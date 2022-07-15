package linked_list

import "fmt"

func mergeListIteration(node1 *LinkedList[int], node2 *LinkedList[int]) *LinkedList[int] {
	var result = &LinkedList[int]{}
	tail := result
	for {
		if node1 == nil {
			tail.Next = node2
			break
		} else if node2 == nil {
			tail.Next = node1
			break
		}
		if node1.Val < node2.Val {
			tail.Next = node1
			node1 = node1.Next
		} else {
			tail.Next = node2
			node2 = node2.Next
		}
		tail = tail.Next
	}
	return result.Next
}

func mergeListRecursion(node1 *LinkedList[int], node2 *LinkedList[int]) *LinkedList[int] {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.Val < node2.Val {
		node1.Next = mergeListRecursion(node1.Next, node2)
		return node1
	} else {
		node2.Next = mergeListRecursion(node1, node2.Next)
		return node2
	}
}

func TestMerge00Iteration() {

	f := LinkedList[int]{Val: 8}
	e := LinkedList[int]{Val: 6, Next: &f}
	d := LinkedList[int]{Val: 5, Next: &e}
	c := LinkedList[int]{Val: 4, Next: &d}
	b := LinkedList[int]{Val: 3, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	two := LinkedList[int]{Val: 7}
	one := LinkedList[int]{Val: 2, Next: &two}

	fmt.Println(mergeListIteration(&a, &one))
}

func TestMerge00Recursion() {

	f := LinkedList[int]{Val: 8}
	e := LinkedList[int]{Val: 6, Next: &f}
	d := LinkedList[int]{Val: 5, Next: &e}
	c := LinkedList[int]{Val: 4, Next: &d}
	b := LinkedList[int]{Val: 3, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	two := LinkedList[int]{Val: 7}
	one := LinkedList[int]{Val: 2, Next: &two}

	fmt.Println(mergeListRecursion(&a, &one))
}

func TestMerge01Iteration() {

	d := LinkedList[int]{Val: 7}
	c := LinkedList[int]{Val: 5, Next: &d}
	b := LinkedList[int]{Val: 3, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	four := LinkedList[int]{Val: 8}
	three := LinkedList[int]{Val: 6, Next: &four}
	two := LinkedList[int]{Val: 4, Next: &three}
	one := LinkedList[int]{Val: 2, Next: &two}

	fmt.Println(mergeListIteration(&a, &one))
}

func TestMerge01Recursion() {

	d := LinkedList[int]{Val: 7}
	c := LinkedList[int]{Val: 5, Next: &d}
	b := LinkedList[int]{Val: 3, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	four := LinkedList[int]{Val: 8}
	three := LinkedList[int]{Val: 6, Next: &four}
	two := LinkedList[int]{Val: 4, Next: &three}
	one := LinkedList[int]{Val: 2, Next: &two}

	fmt.Println(mergeListRecursion(&a, &one))
}

func TestMerge02Iteration() {

	d := LinkedList[int]{Val: 7}
	c := LinkedList[int]{Val: 2, Next: &d}

	eight := LinkedList[int]{Val: 8}
	six := LinkedList[int]{Val: 6, Next: &eight}
	five := LinkedList[int]{Val: 5, Next: &six}
	four := LinkedList[int]{Val: 4, Next: &five}
	three := LinkedList[int]{Val: 3, Next: &four}
	one := LinkedList[int]{Val: 1, Next: &three}

	fmt.Println(mergeListIteration(&c, &one))
}

func TestMerge02Recursion() {

	d := LinkedList[int]{Val: 7}
	c := LinkedList[int]{Val: 2, Next: &d}

	eight := LinkedList[int]{Val: 8}
	six := LinkedList[int]{Val: 6, Next: &eight}
	five := LinkedList[int]{Val: 5, Next: &six}
	four := LinkedList[int]{Val: 4, Next: &five}
	three := LinkedList[int]{Val: 3, Next: &four}
	one := LinkedList[int]{Val: 1, Next: &three}

	fmt.Println(mergeListRecursion(&c, &one))
}

func TestMerge03Iteration() {

	d := LinkedList[int]{Val: 4}
	c := LinkedList[int]{Val: 3, Next: &d}
	b := LinkedList[int]{Val: 2, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	four := LinkedList[int]{Val: 8}
	three := LinkedList[int]{Val: 7, Next: &four}
	two := LinkedList[int]{Val: 6, Next: &three}
	one := LinkedList[int]{Val: 5, Next: &two}

	fmt.Println(mergeListIteration(&a, &one))
}

func TestMerge03Recursion() {

	d := LinkedList[int]{Val: 4}
	c := LinkedList[int]{Val: 3, Next: &d}
	b := LinkedList[int]{Val: 2, Next: &c}
	a := LinkedList[int]{Val: 1, Next: &b}

	four := LinkedList[int]{Val: 8}
	three := LinkedList[int]{Val: 7, Next: &four}
	two := LinkedList[int]{Val: 6, Next: &three}
	one := LinkedList[int]{Val: 5, Next: &two}

	fmt.Println(mergeListRecursion(&a, &one))
}
