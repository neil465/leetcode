func subtractProductAndSum(n int) int {
	cheepe := n
	doobas := []int{}
	for cheepe > 0 {
		doobas = append(doobas, cheepe%10)
		cheepe /= 10
	}
	res1 := 1
	res2 := 0
	for _, i := range doobas {
		res1 *= i
		res2 += i

	}
	return res1 - res2

}