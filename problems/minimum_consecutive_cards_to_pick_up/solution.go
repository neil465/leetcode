func minimumCardPickup(cards []int) int {
    lastIndex := map[int]int{}
    min := 100000
    for i := range cards {
        if _, ok := lastIndex[cards[i]]; ok && min > (i - lastIndex[cards[i]]){
            min = (i - lastIndex[cards[i]])
        }
        lastIndex[cards[i]] = i
    }
    if min == 100000 {
        return -1
    }
    return min + 1
}