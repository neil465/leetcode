var dp = map[int]int{}

func minOperations(nums []int) int {

    m := map[int]int{}
    for _, i := range nums {
        m[i]++
    }
    k := 0
    for _, i := range m {
        g := find(i)
        if g == math.MaxInt32 {
            return -1
        }
        k += g
    }
    return k
}
func find(i int) int {
    if _, ok := dp[i]; ok {
        return dp[i]
    }
    if i ==0 {
        return 0
    }
    if i<0  {
        return math.MaxInt32 
    }
    k := find(i - 2)
    if t := find(i - 3) ; t < k  {
        k= t
    }
    if k == math.MaxInt32  {
        return k
    }
    dp[i] = k + 1

    return k + 1
}