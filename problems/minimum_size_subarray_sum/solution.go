func minSubArrayLen(target int, nums []int) int {
    sum := nums[0]
    i , j := 0, 0
    minLen := 1000000
    for i < len(nums) && j < len(nums) {
        if sum >= target {
            if j - i + 1 < minLen {
                minLen = j - i + 1
            }
            sum -= nums[i]
            i ++
        } else {
            
            j ++
            if j < len(nums) {
                sum += nums[j]
            }
        }
    }
    if minLen == 1000000 {
        return 0
    }
    return minLen
}