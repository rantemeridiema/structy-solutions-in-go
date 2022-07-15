package arrays_and_strings

func Palindrome(firstString, secondString string) bool {
	if len(firstString) != len(secondString) {
		return false
	}
	i, j := 0, len(secondString)-1
	for {
		if firstString[i] != secondString[j] {
			return false
		}
		i++
		j--
		if j < 0 {
			break
		}
	}
	return true
}
