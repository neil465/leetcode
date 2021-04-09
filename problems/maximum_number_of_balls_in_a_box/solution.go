func countBalls(lowLimit int, highLimit int) int {
	mappy := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
        n := i
		for n>0 {
			sum += n%10
			n/=10
		}
		mappy[sum]++
	}
	max := 0
	for _, i2 := range mappy {
		if i2>max {
			max = i2
		}
	}
	return max
}