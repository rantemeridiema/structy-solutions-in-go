package arrays_and_strings

func FiveSort(input []int) []int {
	insertIndex := 0
	for i, elem := range input {
		if elem != 5 {
			temp := input[insertIndex]
			input[insertIndex] = elem
			input[i] = temp
			insertIndex++
		}
	}
	return input
}
