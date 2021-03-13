func sortArrayByParity(A []int) []int {
	res1,res2 := []int{}, []int{}
	for _, i2 := range A {
		if i2%2 == 0{
			res1 = append(res1,i2)
		}else{
			res2 = append(res2,i2)
		}
	}
	res1 = append(res1,res2...)
	return res1
}
