func containsDuplicate(nums []int) bool {
	table := make(map[int]bool)
	for _, num := range nums {
		if _,ok:= table[num];ok {
			return true
		}
		table[num] = true
	}
	return false
}

