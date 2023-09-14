var dp = map[[2]int]int{}

func minDistance(word1 string, word2 string) int {
    dp = map[[2]int]int{}
    return helper(word1, word2, 0, 0)
}
func helper(w1, w2 string, i,j int) int {
    if dp[[2]int{i, j}] > 0 {
        return dp[[2]int{i, j}]
    }
    if len(w1) == i {
        return len(w2) - j
    } 
    if len(w2) == j {
        return len(w1) - i
    }
    if w1[i] == w2[j] {
        return  helper(w1, w2, i + 1, j + 1)
    }
    
    min := 100000
    if temp := 1 + helper(w1, w2, i, j + 1); temp < min {
        min = temp
    }
    if temp := 1 + helper(w1, w2, i + 1, j); temp < min {
        min = temp
    }
    if temp := 1 + helper(w1, w2, i + 1, j + 1); temp < min {
        min = temp
    }
    dp[[2]int{i, j}] = min
    return min

}