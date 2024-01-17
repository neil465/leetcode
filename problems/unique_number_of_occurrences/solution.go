func uniqueOccurrences(arr []int) bool {
    valMap, freqMap := map[int]int{}, map[int]bool{} 

    for _, i := range arr {
        valMap[i] ++
    }
    for _, v := range valMap {
        if freqMap[v]{
            return false
        }
        freqMap[v] = true
    }
    return true
}