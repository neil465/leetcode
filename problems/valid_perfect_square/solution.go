func isPerfectSquare(num int) bool {
	numsquraeroot := 1
	if num == 0 {
		return true
	}

	for numsquraeroot < num+1 {
		if numsquraeroot*numsquraeroot == num {
				
			return true
		}
		numsquraeroot++
	}
	return false

}