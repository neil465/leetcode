func addDigits(num int) int {
	if num < 10 {
		return num
	}
	res := 0

	var x []int
	for num > 0 {
		x = append(x, num%10)
		num = num / 10
	}
	//fmt.Println("array", x)
	for _, i2 := range x {
		res += i2

	}
	//fmt.Println("result", res)

	if res > 9 {
		//fmt.Println("inside the lowest if", res)
		return addDigits(res)

	} else {
		//fmt.Println("resultfromelse", res)
		return res

	}
	return res

}
