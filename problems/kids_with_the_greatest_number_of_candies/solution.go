func kidsWithCandies(candies []int, extraCandies int) []bool {
	arr := []bool{}
	for _, i2 := range candies {
		extra := i2 + extraCandies
        flag := true
		for _, candy := range candies {
            
			if candy > extra {
				arr = append(arr,false)
                flag = false
                break
			}
		}
        if flag{arr = append(arr,true)}
		
        fmt.Println(arr)
	}
	return arr
}