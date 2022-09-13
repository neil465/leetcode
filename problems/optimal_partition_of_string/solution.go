package main

func partitionString(s string) int {
	sum, m := 0, [26]int{}  
	for _, i := range s {
		if m[i - 'a'] > 0  { // if the character is already present in the map
			sum ++ // if the character is already present in the map, then we need to add 1 to the sum
			m = [26]int{} // we reset the map because we are starting a new partition
		}
		m[i - 'a'] ++ // we add the character to the map
	}
	return sum + 1 // we add one, because the last partition is not counted in the loop
}