func singleNumber(nums []int) int {
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        if (i == len(nums) - 1  || nums[i] != nums[i + 1]) && ( i == 0|| nums[i] != nums[i - 1]) {
            return nums[i]
        }

    }
    return -1
}