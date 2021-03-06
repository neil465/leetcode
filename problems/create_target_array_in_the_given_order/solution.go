func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		g := nums[i]
		for j := len(nums) - 2; j >= index[i]; j-- {
			result[j+1] = result[j]
		}
		result[index[i]] = g

	}

	return result
}
