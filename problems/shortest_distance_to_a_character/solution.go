func shortestToChar(s string, c byte) []int {
	result, indexsOfc := []int{}, []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			indexsOfc = append(indexsOfc, i)
		}
	}
	for i := 0; i < len(s); i++ {
		minimum := 100000
		for j := 0; j < len(indexsOfc); j++ {
			if minimum > abs(indexsOfc[j], i) {
				minimum = abs(indexsOfc[j], i)

			}
		}
		result = append(result, minimum)

	}
	fmt.Println(result)
	return result

}
func abs(a int, b int) int {
	if a-b < 0 {
		return -(a - b)
	}
	return a - b
}
