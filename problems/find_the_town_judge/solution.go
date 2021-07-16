func findJudge(n int, trust [][]int) int {
    if len(trust) == 0 && n == 1{
        return 1
    } 
    trusted := make(map[int][]int)
    trusts := make(map[int]bool)
    for _,i := range trust{
        trusts[i[0]] = true
        arr := trusted[i[1]]
        arr = append(arr,0)
        trusted[i[1]] = arr
    }
    for k,v := range trusted{
        if trusts[k] != true && len(v) == n-1{
            return k
        }
    }
    return -1
}