func diagonalPrime(nums [][]int) int {
    isPrime := outputPrimes(4000000)
    max := 0
    fmt.Println(len(nums) == len(nums[0]))
    for i := 0 ; i < len(nums) ; i ++ {

        if (isPrime[nums[i][i]] && nums[i][i] > max){
            max = nums[i][i]
        } 
        if (isPrime[nums[i][len(nums) - 1 - i]] && nums[i][len(nums) - 1 - i] > max){

            max = nums[i][len(nums) - 1 - i]
        } 
    }
    return max
}
func outputPrimes(limit int) map[int]bool {
    primes := map[int]bool{}
	isComposite := make([]bool, limit+1)

	for i := 2; i <= limit; i++ {
		if !isComposite[i] {
			primes[i] = true

			for j := i * i; j <= limit; j += i {
				isComposite[j] = true
			}
		}
	}

	return primes
}