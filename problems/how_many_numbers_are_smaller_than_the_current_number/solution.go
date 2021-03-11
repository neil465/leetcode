func smallerNumbersThanCurrent(nums []int) []int {
	res := []int{}
	for _, num := range nums {
		sum := 0
		
		for _, i2 := range nums {
			if num > i2{
				sum++
			}
		}
		res = append(res,sum)
	}
	return res
}