func buildArray(target []int, n int) []string {
	res := []string{}
	arr := []int{}
	counter := 0
	for i := 1; i < n+1; i++ {
		res = append(res, "Push")
		if target[counter] == i {
            if counter == len(target)-1{
                return res
            }
            if counter < len(target)-1{
                counter++
            }
			arr = append(arr, i)
            
		} else {
			res = append(res, "Pop")
		}
	}
	return res
}