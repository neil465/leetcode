import "sort"

func minimumCost(nums []int) int {
    k := nums[0]
    nums = nums[1:]
    sort.Ints(nums)

    return nums[0] + nums[1] + k
}