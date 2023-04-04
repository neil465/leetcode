func hIndex(citations []int) int {
    finder := map[int]int{}
    maxCount := 0
    for _, i := range citations {
        for j := 0 ; j <= i; j++ {
            finder[j] ++
            if finder[j] == j && j > maxCount {
                maxCount = j
            }
        }
    }
    return maxCount
}