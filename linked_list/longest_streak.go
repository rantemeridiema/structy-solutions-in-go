package linked_list

import "fmt"

type Longest struct {
	resultCount  int
	currentCount int
	prevVal      int
}

func longestStreakIteration(input *LinkedList[int]) int {
	longest := Longest{}
	for input != nil {
		if longest.prevVal == input.Val {
			longest.currentCount++
		} else {
			longest.prevVal = input.Val
			longest.currentCount = 1
		}
		if longest.currentCount > longest.resultCount {
			longest.resultCount = longest.currentCount
		}
		input = input.Next
	}
	return longest.resultCount
}

func longestStreakRecursion(input *LinkedList[int], longest Longest) int {
	if input == nil {
		return longest.resultCount
	}
	if input.Val == longest.prevVal {
		longest.currentCount = longest.currentCount + 1
	} else {
		longest.currentCount = 1
		longest.prevVal = input.Val
	}
	if longest.currentCount > longest.resultCount {
		longest.resultCount = longest.currentCount
	}
	return longestStreakRecursion(input.Next, longest)
}

func TestLongestStreak00() {
	test := NewLinkedList([]int{5, 5, 7, 7, 7, 6})
	fmt.Println(test, " : ", longestStreakIteration(test))
	fmt.Println(test, " : ", longestStreakRecursion(test, Longest{}))
}

func TestLongestStreak01() {
	test := NewLinkedList([]int{3, 3, 3, 3, 9, 9})
	fmt.Println(test, " : ", longestStreakIteration(test))
	fmt.Println(test, " : ", longestStreakRecursion(test, Longest{}))
}

func TestLongestStreak02() {
	test := NewLinkedList([]int{9, 9, 1, 9, 9, 9})
	fmt.Println(test, " : ", longestStreakIteration(test))
	fmt.Println(test, " : ", longestStreakRecursion(test, Longest{}))
}

func TestLongestStreak03() {
	test := NewLinkedList([]int{5, 5})
	fmt.Println(test, " : ", longestStreakIteration(test))
	fmt.Println(test, " : ", longestStreakRecursion(test, Longest{}))
}

func TestLongestStreak04() {
	test := NewLinkedList([]int{4})
	fmt.Println(test, " : ", longestStreakIteration(test))
	fmt.Println(test, " : ", longestStreakRecursion(test, Longest{}))
}

func TestLongestStreak05() {
	fmt.Println(nil, " : ", longestStreakIteration(nil))
	fmt.Println(nil, " : ", longestStreakRecursion(nil, Longest{}))
}
