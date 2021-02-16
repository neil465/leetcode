func isPowerOfFour(n int) bool {
	powersum := 1
	for powersum<=n {
		if powersum == n {
			return true
		}
		powersum*= 4
	}
	return false
}
