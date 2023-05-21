

func punishmentNumber(n int) int {
    v := 0
    for i := 1 ; i <= n ; i ++ {
        if find(0, i, splitInt(i * i)) {
           v +=  i * i
        } 
    }
    return v
}
func find(sum, g int, digits []int) bool {
  
    if sum == g && len(digits) == 0 {

        return true
    }
    for i := 1 ; i <= len(digits); i++ {
        v ,_ := strconv.Atoi(arrayToString(digits[:i ], ""))
        if find(sum + v, g, digits[i :]) {
            return true
        }
    }
    return false
    
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
        slc = append([]int{n%10}, slc...)
		n = n / 10
	}
	return slc
}
func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]") 
}