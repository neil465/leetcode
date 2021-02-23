func checkPerfectNumber(num int) bool {
	res := 0
	for i := 1; i < num/2 +1; i++ {
		if num%i == 0 {
			res += i
		}
	}
	return res == num
}
