func sortArrayByParityII(nums []int) []int {

	odds, evens,res := []int{}, []int{},[]int{}
	for _, num := range nums {
		if num%2 == 0 {
			evens = append(evens, num)
		} else {
			odds = append(odds, num)
		}
	}
	for i := 0; i < len(odds); i++ {
		res = append(res,evens[i])
		res = append(res,odds[i])
	}
	return res
}