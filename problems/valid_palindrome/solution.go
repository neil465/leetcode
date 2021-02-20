func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	splits := []rune(s)
	for i, split := range splits {
		if !unicode.IsLetter(split) && !unicode.IsDigit(split) {
			splits[i] = -1
		}
	}
	for left < right {
		for left <= len(s)-1 && left < right && splits[left] == -1 {
			left++
		}
		for right >= 0 && right > left && splits[right] == -1 {
			right--
		}
		if strings.ToLower(string(splits[left])) != strings.ToLower(string(splits[right])) {
			return false
		}
		left++
		right--
	}
	return true
}