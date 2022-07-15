package arrays_and_strings

import (
	"fmt"
	"strconv"
)

//https://structy.net/problems/premium/uncompress
//input : 2c3a
//output : ccaaa

func Uncompress(input string) {
	firstIndex := 0
	var result string
	for i := 1; i < len(input); i++ {
		currentChar := string(input[i])
		if _, err := strconv.Atoi(currentChar); err != nil {
			length, _ := strconv.Atoi(input[firstIndex:i])
			result = result + repeat(length, currentChar)
			firstIndex = i + 1
		}
	}
	fmt.Println(result)
}

func repeat(count int, char string) string {
	var result string
	for i := 1; i <= count; i++ {
		result = result + char
	}
	return result
}
