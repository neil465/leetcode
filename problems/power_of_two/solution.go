func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }
	flag := false
	for n > 0 {
		if n&1 == 1 {
			if flag {
				return false
			}
			flag = true
		}
		n = n >> 1
	}
	return true
}
