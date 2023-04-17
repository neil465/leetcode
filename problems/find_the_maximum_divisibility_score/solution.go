func maxDivScore(nums []int, divisors []int) int {
    max, v := 0,divisors[0]
    for _, divisor := range divisors {
        t := 0
        for _, num := range nums {
            if num % divisor == 0 {
                t++
            }
        }
        if t > max {
            max = t
            v = divisor
        }
        if t == max && v > divisor {
            v = divisor 
        }
    }
    return v
}