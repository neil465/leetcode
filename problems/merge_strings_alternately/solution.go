func mergeAlternately(word1 string, word2 string) string {
	mergestringres := ""
	for i, i2 := range word1 {
		mergestringres += string(i2)
		
		if len(word2) > i {
			mergestringres += string(word2[i])
			
		}
	}
	if len(word1) < len(word2) {
		mergestringres += word2[len(word1):]
	}
	//fmt.Println(mergestringres)
	return mergestringres
}