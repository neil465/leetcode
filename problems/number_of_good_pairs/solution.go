func numIdenticalPairs(nums []int) int {
	result:= 0
	table := make(map[int]int)
	for _, num := range nums {
		table[num] += 1
	}
	for i2, i := range table {
		table[i2] = allposibilitys(i)
	}
	for _, i2 := range table {
		result+=i2
	}
	return result
}
func allposibilitys(num int) int {
	return ((num * (num+1))/2)-num
}
