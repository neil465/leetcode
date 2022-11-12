func averageValue(nums []int) int {
    sum, times := 0, 0
    for _, i := range nums {
        if i % 6 == 0 {
            sum, times = sum + i, times + 1
        }
    }
    
    if times == 0 { return 0 }
    
    return sum / times
}