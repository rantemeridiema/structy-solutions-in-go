package arrays_and_strings

//https://www.structy.net/problems/premium/intersection
//input : [4,2,1,6], [3,6,9,2,10]
//output [2,6]
//constraint: each array will not contain duplicate

func Intersection(i1 []int, i2 []int) []int {
	dict := make(map[int]bool)
	var result []int
	for _, elem := range i2 {
		dict[elem] = true
	}
	for _, elem := range i1 {
		if _, exist := dict[elem]; exist {
			result = append(result, elem)
		}
	}

	return result
}
