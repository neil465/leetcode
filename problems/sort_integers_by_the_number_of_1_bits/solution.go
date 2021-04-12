func sortByBits(arr []int) []int {
	res := []int{}
	mappy := make(map[int][]int)
	for _, i2 := range arr {
		num := i2
		counter := 0
		for num > 0 {
			if (num & 1) == 1 {
				counter++
			}
			num >>= 1
		}
		mappy[counter] = append(mappy[counter], i2)
	}
	for i := 0; i < len(mappy); i++ {
		k := mappy[i]
		sort.Ints(k)
		mappy[i] = k
		res = append(res, mappy[i]...)
	}
	return res
}
