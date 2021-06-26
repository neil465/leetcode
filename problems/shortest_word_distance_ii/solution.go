type WordDistance struct {
    k []string
}


func Constructor(wordsDict []string) WordDistance {
    return WordDistance{k:wordsDict}
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    
    min := 999999999909099090
    w1 := -1
    w2 := -1
    
    for i := 0 ; i < len(this.k) ; i++{
        if this.k[i] == word1{
            w1 = i
            if w2 != -1 && abs(w2 - w1) < min{
                min = abs(w2 - w1)
            }
        }
        if this.k[i] == word2{
            w2 = i
            if w1 != -1 && abs(w1 - w2) < min{
                min = abs(w1 - w2)
            }
        }
    }
    return min
}


func abs (i int)int {
    if i < 0 {
        return i*-1
    }
    return i
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */