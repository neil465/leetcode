func selfDividingNumbers(left int, right int) []int {

	resarray := []int{}
	for num := left; num <= right; num++ {

		if num < 10 {
			resarray = append(resarray, num)
		} else {
			b := num
			arr := []int{}
			for b > 0 {
				arr = append(arr, b%10)
				b /= 10
			}

			flag := true
			for _, digit := range arr {
				//fmt.Println("inside the for loop in range of arr")
				if digit != 0 {
					if num%digit != 0 {
						//fmt.Println()
						flag = false
						break
					}
                    
                }else{
                    flag = false
                    break
                }

			}
			if flag {
				resarray = append(resarray, num)
			}
		}
		//fmt.Println("restult in proggres", resarray)

	}
	

	return resarray
}
