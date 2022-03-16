var result [][]int

func subsets(nums []int) [][]int {
    result = [][]int{}
    backtrack(nums, []int{})
    return result
}
func backtrack(nums []int, res []int) {
    
    result = append(result, append([]int{}, res...))
    
    for i := range nums{
        backtrack(nums[i + 1:], append(res, nums[i]))
    }
}