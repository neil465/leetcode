
func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := map[float64]bool{}
	for len(nums) > 0 {
        
		m[float64(nums[0]+nums[len(nums)-1]) / float64(2)] = true
		nums = nums[1:len(nums) - 1]
	}
	return len(m)
}