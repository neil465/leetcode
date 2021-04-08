func uniqueOccurrences(arr []int) bool {
	mappy := make(map[int]int)
	arr2 := []int{}
	for _, i2 := range arr {
		mappy[i2]++
	}
	for _, i2 := range mappy {
        
		for i := 0; i < len(arr2); i++ {
			if i2 == arr2[i] {
				return false
			}
            
		}
        arr2 = append(arr2, i2)
			
	}
	return true
}