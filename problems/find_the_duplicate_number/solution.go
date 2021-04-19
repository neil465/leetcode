import "sort"
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i]  {
			return nums[i]
		}
	}
	return 0
}