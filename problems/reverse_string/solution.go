func reverseString(s []byte) {

	right := len(s)-1
	left := 0
	for right > left {
		s[right], s[left] = s[left], s[right]
		right --
		left ++
	}
}
