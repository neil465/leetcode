func singleNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] != nums[i] && nums[i+1] != nums[i] {
			return nums[i]
		}
	}
	return 0
}
