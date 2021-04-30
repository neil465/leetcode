
func findPeakElement(nums []int) int {
	greatest := [2]int{nums[0],0}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= greatest[0] {
			greatest[0] = nums[i]
			greatest[1] = i
            fmt.Println(greatest)
		}
	}
	return greatest[1]
}
