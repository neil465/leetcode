
func countPrimes(n int) int {
	result := 0
    res := []int{}
	for i := 2; i < n; i++ {
		
		flag := true
		for j := 2; j * j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
            res = append(res,i)
			result++
		}
		
	}
    fmt.Println(res)
	return result
}