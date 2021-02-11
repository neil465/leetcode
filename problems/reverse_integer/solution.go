// import (
// 	"fmt"
// 	"strconv"
// )
// func reverse(x int) int {
// 	b := 0
// 	if x < 0 {

// 		b = x
// 		x *= -1

// 	}
// 	if x == 0 {
// 		return 0
// 	}
// 	i := []int{}
// 	for x > 0 {
// 		i = append(i, x%10)
// 		x = x / 10
// 	}
// 	res := ""
// 	for _, i3 := range i {

// 		res += fmt.Sprint(i3)

// 	}
// 	if b < 0 {
// 		res = "-" + res

// 	}

// 	n, _ := strconv.Atoi(res)
// 	if n < -2147483648 {
// 		return 0
// 	}
// 	if n > 2147483647 {
// 		return 0
// 	}
// 	return n
// }
package main

import (
	"fmt"
	"strconv"
)


func reverse(x int) int {
	if x == 0 {
		return 0
	}

	original := 0
	if x < 0 {
		original = x
		x *= -1
	}

	ints := []int{}
	for x > 0 {
		ints = append(ints, x%10)
		x = x / 10
	}

	res := ""
	for _, i3 := range ints {
		res += fmt.Sprint(i3)

	}
	if original < 0 {
		res = "-" + res
	}

	n, _ := strconv.Atoi(res)

	if n < -2147483648 {
		return 0
	}
	
	if n > 2147483647 {
		return 0
	}

	return n
}

