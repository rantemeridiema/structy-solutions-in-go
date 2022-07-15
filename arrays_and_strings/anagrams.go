package arrays_and_strings

func Anagrams(a, b string) bool {
	charMap := make(map[rune]int)
	for _, char := range a {
		charMap[char] = charMap[char] + 1
	}
	for _, char := range b {
		charMap[char] = charMap[char] - 1
		if charMap[char] < 0 {
			return false
		} else if charMap[char] == 0 {
			delete(charMap, char)
		}
	}
	if len(charMap) > 0 {
		return false
	}
	return true
}
