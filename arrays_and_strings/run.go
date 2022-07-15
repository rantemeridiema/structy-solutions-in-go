package arrays_and_strings

import "fmt"

func Run() {
	//Uncompress
	Uncompress("3c2a")
	Uncompress("3n12e2z")

	//Compress
	Compress("aaaxxzzzz")
	Compress("bcccccee")

	//Palindrome
	fmt.Println(Palindrome("paper", "repap"))
	fmt.Println(Palindrome("paper", "somet"))
	fmt.Println(Palindrome("paper", "hing else"))

	//Anagrams
	fmt.Println(Anagrams("paper", "repap"))
	fmt.Println(Anagrams("cats", "tocs"))
	fmt.Println(Anagrams("monkeyswrite", "newyorktimes"))
	fmt.Println(Anagrams("tax", "taxi"))
	fmt.Println(Anagrams("abbc", "aabc"))
	fmt.Println(Anagrams("restful", "fluster"))

	//Most Frequent Char
	fmt.Println(GetMostFrequent("bookkeeper"))
	fmt.Println(GetMostFrequent("david"))
	fmt.Println(GetMostFrequent("jeepoor"))
	fmt.Println(GetMostFrequent("mississippi"))
	fmt.Println(GetMostFrequent("potato"))
	fmt.Println(GetMostFrequent("abby"))
	fmt.Println(GetMostFrequent("eleventennine"))
	fmt.Println(GetMostFrequent("riverbed"))

	//Pair Sum
	fmt.Println(PairSum([]int{3, 2, 5, 4, 1}, 8))
	fmt.Println(PairSum([]int{4, 7, 9, 2, 5, 1}, 5))
	fmt.Println(PairSum([]int{4, 7, 9, 2, 5, 1}, 3))
	fmt.Println(PairSum([]int{1, 6, 7, 2}, 13))
	fmt.Println(PairSum([]int{9, 9}, 18))
	fmt.Println(PairSum([]int{6, 4, 2, 8}, 12))

	//Pair Product
	fmt.Println(PairProduct([]int{3, 2, 5, 4, 1}, 8))
	fmt.Println(PairProduct([]int{3, 2, 5, 4, 1}, 10))
	fmt.Println(PairProduct([]int{4, 7, 9, 2, 5, 1}, 5))
	fmt.Println(PairProduct([]int{4, 7, 9, 2, 5, 1}, 35))
	fmt.Println(PairProduct([]int{4, 6, 8, 2}, 16))

	//Intersection
	fmt.Println(Intersection([]int{4, 2, 1, 6}, []int{3, 6, 9, 2, 10}))
	fmt.Println(Intersection([]int{2, 4, 6}, []int{4, 2}))
	fmt.Println(Intersection([]int{4, 2, 1}, []int{1, 2, 4, 6}))
	fmt.Println(Intersection([]int{0, 1, 2}, []int{10, 11}))
	a := make([]int, 10)
	b := make([]int, 10)
	for i := 1; i <= 10; i++ {
		a[i-1] = i
		b[i-1] = i
	}
	fmt.Println(Intersection(a, b))

	//Five Sort
	fmt.Println(FiveSort([]int{12, 5, 1, 5, 12, 7}))
	fmt.Println(FiveSort([]int{5, 2, 5, 6, 5, 1, 10, 2, 5, 5}))
}
