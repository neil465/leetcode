func countWays(nums []int) int {
    sort.Ints(nums)
    j := 0
    if nums[0] > 0 {
        j++
    }
    for i := 0 ; i < len(nums); i ++ {
        if nums[i] < i + 1 && (i == len(nums) -1 || nums[i + 1] > i + 1){

            j ++
        }
    }
    return j
}

