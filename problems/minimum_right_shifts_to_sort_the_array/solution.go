func minimumRightShifts(nums []int) int {
    for i := 0 ; i < len(nums); i++ {
        k := true
        for j := 0; j < len(nums) - 1; j++ {
            if nums[j] > nums[j + 1] {
                k = false
                break
            }
        }
        if k { return i }
        nums = append([]int{nums[len(nums) - 1]}, nums[:len(nums) - 1]...)
    }
    return -1
}