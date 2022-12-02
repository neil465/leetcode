func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
    e1, e2, f1, f2 := [26]bool{}, [26]bool{}, []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    for i := range word1 {
        e1[word1[i] - 'a'] = true
        e2[word2[i] - 'a'] = true
        f1[word1[i] - 'a'] ++
        f2[word2[i] - 'a'] ++
    }
    sort.Ints(f1)
    sort.Ints(f2)
    for i := range e1 {
        if e1[i] != e2[i] {
            return false
        }
    }
    for i := range f1 {
        if f1[i] != f2[i] {
            return false
        }
    }
    
    return true
	
	
}