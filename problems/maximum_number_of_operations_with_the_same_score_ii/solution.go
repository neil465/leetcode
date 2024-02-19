
import "math"

var dp = map[[3]int]int{}
func maxOperations(nums []int) int {
    dp = map[[3]int]int{}
    return re(-1, 0, len(nums) - 1, nums)
}
func re(wantscore int , l, r int, nums []int) int {
    if val, ok := dp[[3]int{wantscore, l, r}]; ok {
        return val
    }
    if l == len(nums) || r == -1 || l >= r {
        return 0
    }
    cur := math.MinInt32
    if wantscore == -1 || (len(nums) - l > 2 && nums[l] + nums[l + 1] == wantscore) {
        cur = max(cur, re(nums[l] + nums[l + 1], l + 2, r, nums) + 1)
    }
    if wantscore == -1 || (r > 1 && nums[r] + nums[r - 1] == wantscore) {
        cur = max(cur, re(nums[r] + nums[r - 1], l, r - 2, nums) + 1)
    }
    if wantscore == -1 ||  nums[l] + nums[r ] == wantscore {
        cur = max(cur, re(nums[l] + nums[r ], l + 1, r - 1, nums) + 1)
    }
    if cur == math.MinInt32 {
        return 0
    }
    dp[[3]int{wantscore, l, r}] = cur
    return cur
}