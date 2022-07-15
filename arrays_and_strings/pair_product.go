package arrays_and_strings

func PairProduct(input []int, target int) []int {
	complements := make(map[int]int)
	for index, current := range input {
		modulo := target % current
		if modulo != 0 {
			continue
		}
		complement := target / current
		if _, isExist := complements[complement]; isExist {
			return []int{complements[complement], index}
		} else {
			complements[current] = index
		}
	}
	return nil
}
