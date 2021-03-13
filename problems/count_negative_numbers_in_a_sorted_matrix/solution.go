func countNegatives(grid [][]int) int {
	res := 0
	for _, ints := range grid {
		for _, i2 := range ints {
			if i2<0 {
				res++
			}
		}
	}
	return res
}