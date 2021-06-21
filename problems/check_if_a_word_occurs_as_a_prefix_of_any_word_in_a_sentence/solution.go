func isPrefixOfWord(sentence string, searchWord string) int {
    g := strings.Split(sentence," ")
    for j,i := range g {
        if len(i) >= len(searchWord){
            if i[:len(searchWord)] == searchWord{
		        return j+1
	        }   
        }
    }
    return -1
}