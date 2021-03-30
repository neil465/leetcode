func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	res := ""
	res2 := ""
	for i := 0; i < len(word1); i++ {
		res += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		res2 += word2[i]
	}
	if len(res) != len(res2) {
		return false
	}
	for i := 0; i < len(res); i++ {
		if res[i] != res2[i] {
			return false
		}
	}
	return true
}