func numberOfMatches(n int) int {
	res := 0
	for n!=1 {
		res+=n/2
		n = n-(n/2)
	}
	return res
}