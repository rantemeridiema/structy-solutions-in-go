package arrays_and_strings

func PairSum(input []int, target int) []int {
	complements := make(map[int]int)
	for index, current := range input {
		discrepancy := target - current
		if _, isExist := complements[discrepancy]; isExist {
			return []int{complements[discrepancy], index}
		} else {
			complements[current] = index
		}
	}
	return nil
}
