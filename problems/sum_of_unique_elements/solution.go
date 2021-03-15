func sumOfUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] || i != 0 && nums[i] == nums[i-1] {
			continue
		} else {
			res += nums[i]
		}
	}
	return res
}
