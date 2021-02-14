func wordPattern(pattern string, s string) bool {
	patternTable := make(map[string]string)
	wordTable := make(map[string]string)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, w := range words {

		word, doesItExists := patternTable[string(pattern[i])]
		char, doesItExists1 := wordTable[w]

		if doesItExists {
			if w != word {
				return false
			}
		} else {
			patternTable[string(pattern[i])] = w

		}

		//fmt.Println("hello", doesItExists1, w)
		if doesItExists1 {

			if string(pattern[i]) != char {
				return false
			}
		} else {
			wordTable[w] = string(pattern[i])
		}
		fmt.Println(patternTable, wordTable)

	}

	return true

}
