func numJewelsInStones(jewels string, stones string) int {
	result := 0
	for _, i := range jewels {
		for _, i2 := range stones {
			if i==i2{
				result++
			}
		}
	}
	return result

}