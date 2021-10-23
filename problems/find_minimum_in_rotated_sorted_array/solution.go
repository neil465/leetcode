func findMin(nums []int) int {
    min := 100000
    l,r := 0,len(nums)-1 
    for l <=r {
        if nums[l] < min{ min = nums[l]}
        if nums[r] < min { min = nums[r]}
        l++
        r--
    }
    return min
}