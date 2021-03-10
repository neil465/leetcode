func flipAndInvertImage(image [][]int) [][]int {
	res := [][]int{}
	for _, ints := range image {
		A := []int{}
		A = invert(ints)
		A = flip(ints)
		res = append(res, A)
	}
	return res
}
func flip(arr []int) []int {
	res := []int{}
	for _, i2 := range arr {
		res = append([]int{i2}, res...)
	}
	return res
}
func invert(arr []int) []int {
	for i, i2 := range arr {
		if i2 == 0 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}
	}
	return arr
}
