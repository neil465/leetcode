func returnToBoundaryCount(nums []int) int {
    cur := 0 
    count := 0

    for _, i := range nums {
        cur += i
        if cur == 0 {
            count ++
        }
    }
    return count
}