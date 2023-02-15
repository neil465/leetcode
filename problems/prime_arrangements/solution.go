func numPrimeArrangements(n int) int {
    sum := 1
    sum2 := 1
    primes := countPrimes(n + 1)
    fmt.Println(n - primes)
    for i := 1 ; i <= n - primes; i++ {
        sum *= i
        sum = sum % 1000000007
        fmt.Println(sum)
    }
    fmt.Println(sum)
    for i := 1 ; i <= primes; i++ {
        sum2 *= i
        sum2 = sum2 % 1000000007
    }
    fmt.Println(sum, sum2)
    return sum * sum2 % 1000000007
}
func countPrimes(n int) int {
	result := 0
    res := []int{}
	for i := 2; i < n; i++ {
		
		flag := true
		for j := 2; j * j <= i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
            res = append(res,i)
			result++
		}
		
	}
    fmt.Println(res)
	return result
}
