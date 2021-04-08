func majorityElement(nums []int) []int {
    res := []int{}
	mappy := make(map[int]int)
	for _, num := range nums {
		mappy[num]++
	}
	for i, i2 := range mappy {
		if i2>len(nums)/3{
            res = append(res,i)
		}
	}
	return res
}