func findKOr(nums []int, k int) int {
    sum := 0
    for i := 0; i < 39 ; i ++ { 
        v := 0 
        check := 1
        for j := 0 ; j < i; j ++ {
            check *= 2
        }
        for _, value := range nums {
            if value & check != 0 {
                v ++
            }
        }
        // fmt.Println(v, check)
        if v >= k {
            
            add := 1
            for times := 0 ; times < i; times ++ {
                add *= 2
            }
            sum += add
        }
    }
    return sum
}
