func majorityElement(nums []int) int {
	mappy := make(map[int]int)
	for _, num := range nums {
		mappy[num]++
	}
	for i, i2 := range mappy {
		if i2>len(nums)/2 {
			return i
		}
	}
	return 1
}