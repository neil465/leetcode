func isUgly(num int) bool {
	if num == 0{
		return false
	}
	isDivisable := true
	for isDivisable {
		isDivisable = false
		if num%2 == 0 {
			isDivisable = true
			num /= 2
		}
		if num%3 == 0 {
			num /= 3
			isDivisable = true
		}
		if num%5 == 0 {
			num /= 5
			isDivisable = true
		}
	}
	return num == 1
}
