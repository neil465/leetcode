func countHillValley(nums []int) int {
    c := 0
    last := -1
    for i := 0; i < len(nums) - 1; i++ {
        if last != -1 && ((nums[i] < nums[i + 1] && nums[i] < last) || (nums[i] > nums[i + 1] && nums[i] > last)) {
            c++
        } 
        if nums[i] != nums[i + 1] {
            last = nums[i]
        }
    }
    return c
}