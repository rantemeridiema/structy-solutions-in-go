package arrays_and_strings

import (
	"fmt"
	"strconv"
)

//https://structy.net/problems/premium/compress
//input : 'aaazzxxxx'
//output: 3a2z4x

func Compress(input string) {
	var result string
	currentChar := input[0]
	count := 1
	for i := 1; i < len(input); i++ {
		if input[i] != currentChar {
			result = result + getCount(count) + string(currentChar)
			currentChar = input[i]
			count = 0
		}
		count++
	}
	result = result + getCount(count) + string(currentChar)
	fmt.Println("Result for ", input, " is ", result)
}

func getCount(count int) string {
	if count == 1 {
		return ""
	}
	return strconv.Itoa(count)
}
