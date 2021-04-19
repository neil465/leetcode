func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			nums = append(nums[:i-1], nums[i-1+1:]...)
			i--
		}
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+1+1:]...)
		}
	}
	return len(nums)

}
