package arrays_and_strings

import "fmt"

//https://structy.net/problems/premium/most-frequent-char
//input : bookeeper
//output : e
// if the count is equal, return the char that appear first

func GetMostFrequent(input string) string {
	chars := make(map[rune][]int)
	max := 0
	maxIndex := 0
	for pos, char := range input {
		value, isExist := chars[char]

		if isExist {
			value[0] = chars[char][0] + 1
		} else {
			chars[char] = append(value, 1, pos)
		}

		count := chars[char][0]
		currentFirstIndex := chars[char][1]
		if count > max || (count == max && currentFirstIndex < maxIndex) {
			max = count
			maxIndex = currentFirstIndex
		}
	}
	return fmt.Sprint("Most frequent char of ", input, " is ", string(input[maxIndex]))
}
