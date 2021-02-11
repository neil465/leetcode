import (
	"fmt"
	"strconv"
)
func isPalindrome(x int) bool {
	b := x
	if x < 0 {
		return false
	}
	i := []int{}
	for x > 0 {
		i = append(i, x%10)
		x /= 10
	}

	res := ""
	for _, i3 := range i {
		res += fmt.Sprint(i3)

	}
	
	n, _ := strconv.Atoi(res)


	if n == b {
		return true
	}
	return false
}
