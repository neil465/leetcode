func hammingWeight(num uint32) int {
	res := 0
	for int(num)>0 {
		if num & 1 == 1 {
			res += 1
		}
		num>>=1
	}
	return res
}