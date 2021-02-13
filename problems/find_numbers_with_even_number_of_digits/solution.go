func findNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		i := 0
		for num>0 {
			i++
			num/=10
		}
		if i%2 == 0 {
			res++
		}
	}
	return res

}
