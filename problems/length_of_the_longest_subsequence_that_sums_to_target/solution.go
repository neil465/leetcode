func lengthOfLongestSubsequence(nums []int, target int) int {
    return re(nums, target, 0,0, map[[2]int]int{})
}
func re(nums []int, target, cursum, curind int, dp map[[2]int]int) int {
    if cursum == target {
        return 0
    } 
    if cursum > target {
        return -1
    }
    if curind == len(nums) {
        return -1
    }
    if val, ok := dp[[2]int{curind , cursum}]; ok {
        return val
    }
    v1, v2 := re(nums, target, cursum + nums[curind], curind + 1, dp) + 1, re(nums, target, cursum, curind + 1, dp)
    if v1 + 1> v2 && v1 != 0{
        dp[[2]int{curind, cursum}] = v1 
        return v1 
    } else {
        dp[[2]int{curind, cursum}] = v2
        return v2
    }
    return -1

    
}