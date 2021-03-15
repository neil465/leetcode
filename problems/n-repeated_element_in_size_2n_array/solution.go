func repeatedNTimes(A []int) int {
	mappy := make(map[int]int)
	for _, i2 := range A {
		mappy[i2]++
	}
	max:=0
	key :=0
	for i, i2 := range mappy {
		if i2>max {
			max = i2
			key = i
		}
	}
	return key
}