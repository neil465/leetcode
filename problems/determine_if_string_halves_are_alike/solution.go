func halvesAreAlike(s string) bool {
	vouls := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	sum := 0
	sum2 := 0
	firstHalf := s[:len(s)/2]
	lastHalf := s[len(s)/2:]
	for i, i2 := range firstHalf {
		for _, voul := range vouls {
			if voul == string(i2) {
				sum++
			}
			if voul == string(lastHalf[i]) {
				sum2++
			}
		}
	}
	return sum == sum2
}
