func maximumWealth(accounts [][]int) int {
	max := -1
	for _, account := range accounts {
		res := 0
		for _, inty := range account {
			res += inty
		}
		if res>max {
			max = res
		}
		
	}
	return max
}
