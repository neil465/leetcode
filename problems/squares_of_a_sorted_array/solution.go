func sortedSquares(nums []int) []int {
	result := []int{}
	left := 0
	right := len(nums) - 1
	for left <= right {
        l,r := nums[left]*nums[left] , nums[right]*nums[right]
		if l > r {
			result = append([]int{l} , result...)
			left++
		}else{
			result = append([]int{r} , result...)
			right--
		}
	}
	return result
}