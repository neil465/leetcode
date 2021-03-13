func findMaxAverage(nums []int, k int) float64 {
	res := 0
	for i := 0; i < k; i++ {
		res += nums[i]
	}
	max := res
	counter := k
	for counter < len(nums) {
		res = res - nums[counter-k] + nums[counter]
		if res > max {
			max = res
		}
		counter++
	}
	return float64(max) / float64(k)
}
