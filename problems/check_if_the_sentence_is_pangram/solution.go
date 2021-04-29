func checkIfPangram(sentence string) bool {
	mappy := make([]int,26)
	
	for _, i2 := range sentence {
		mappy[i2 - 'a']++
	}
    fmt.Println(mappy)
	for _,v := range mappy {
        if v == 0 {
			return false
		}
	}
	return true
}