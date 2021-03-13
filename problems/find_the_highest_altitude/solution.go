func largestAltitude(gain []int) int {
	max := 0
	altitude :=0
	for _, i2 := range gain {
		altitude += i2
		if altitude > max{
			max = altitude
		}
	}
	return max
}