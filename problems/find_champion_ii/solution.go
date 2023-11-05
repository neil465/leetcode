func findChampion(n int, edges [][]int) int {
    m := map[int]bool{}
    for _, i := range edges {
        m[i[1]] = true
    }
    cnt := 0 
    ind := 0 
    for i := 0 ; i < n ; i ++ {
        if !m[i] {
            cnt ++
            ind = i
        }
    }
    if cnt == 1{
        return ind
    }
    return -1 
}