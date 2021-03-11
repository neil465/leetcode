func countGoodRectangles(rectangles [][]int) int {
	max := 0
	counter := 0
	mappy := make(map[int]int)
	for _, rectangle := range rectangles {
		mappy[min(rectangle[0],rectangle[1])]++
	}
	for i, i2 := range mappy {
		if i>max {
			max = i
			counter = i2
		}
	}
	return counter
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}