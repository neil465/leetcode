func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
        if i == len(nums)-1 && target > num{
            return i+1
        }
        if i < len(nums) && num < target && nums[i+1]> target{
            return i+1
        }
        
	}
	
	return 0
}
