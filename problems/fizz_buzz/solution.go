import "fmt"

func fizzBuzz(n int) []string {
	arr := []string{}
	for i :=1 ; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0  {
			arr = append(arr, "Buzz")
		}else {
			arr = append(arr, fmt.Sprint(i))
		}
		
	}
	return arr
}