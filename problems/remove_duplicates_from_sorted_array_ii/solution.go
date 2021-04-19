func removeDuplicates(nums []int) int {
	mappy := make(map[int]int)
	for _, num := range nums {
		mappy[num]++
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if mappy[nums[j]] > 2 {
				mappy[nums[j]]--
				nums = append(nums[:j], nums[j+1:]...)
			}
		}
	}
	return len(nums)
}