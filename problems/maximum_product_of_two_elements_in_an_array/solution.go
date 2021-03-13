func maxProduct(nums []int) int {
	 max1 , max2 := 0,0
	 index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >max1 {
			max1 = nums[i]
			index = i
		}
	}
	nums = append(nums[:index], nums[index+1:]...)
	for i := 0; i < len(nums); i++ {
		if nums[i] >max2 {
			max2 = nums[i]
			
		}
	}
	return (max1-1)*(max2-1)
	
}