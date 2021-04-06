func moveZeroes(nums []int) {
	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[counter] = nums[counter], nums[i]
			counter++
		}
	}
}
