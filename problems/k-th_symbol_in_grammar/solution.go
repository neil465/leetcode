func kthGrammar(n int, k int) int {
    if n == 1 {
        return 0
    }
    if k % 2 == 0  {
        return flip(kthGrammar(n - 1, (k + 1) / 2))
    } 
    return kthGrammar(n - 1, (k + 1) / 2)
}
func flip( i int) int {
    if i == 0 {
        return 1
    }
    return 0
}