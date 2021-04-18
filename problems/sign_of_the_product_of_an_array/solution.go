func arraySign(nums []int) int {
	sign := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0{
			if sign == 0 {
				sign = -1
			}else{
				sign = 0
			}
		}
		
	}
	if sign == 0 {
		return 1
	}else {
		return -1
	}
}