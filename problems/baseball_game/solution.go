import "strconv"

func calPoints(ops []string) int {
	res := []int{}
	for i := 0; i < len(ops); i++ {
		if _, err := strconv.Atoi(ops[i]); err == nil {
			n, _ := strconv.Atoi(ops[i])
			res = append(res, n)
		}
		if ops[i] == "C" {
            res = append(res[:len(res)-1], res[len(res):]...)
		}
		if ops[i] == "D" {
            fmt.Println(res)
            res = append(res, res[len(res)-1] * 2)
		}
		if ops[i] == "+" {
            res = append(res, res[len(res)-1] + res[len(res)-2])
		}
	}
	result := 0
	for _, re := range res {
		result += re
	}
	return result
}
