func numEquivDominoPairs(dominoes [][]int) int {
    res := 0
    for j,i := range dominoes {
        sort.Ints(i)
        dominoes[j] = i
    }
    for k,i := range dominoes{
        for j:= k+1 ; j< len(dominoes) ; j++{
            if i[0] == dominoes[j][0] && i[1] == dominoes[j][1]{
                
                res++
            }
        }
    }
    return res
}