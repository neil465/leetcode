func getMinDistance(nums []int, target int, start int) int {
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			arr = append(arr, i)
		}
	}
	minimum := 100000
	for _, i2 := range arr {
		if absolute(i2,start)<minimum {
			minimum = absolute(i2,start)
		}
	}
	return minimum
}
func absolute(a int, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	}
	return a - b
}