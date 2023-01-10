func distinctPrimeFactors(nums []int) int {
    a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

    m := map[int]bool{}

    for _, i := range nums {


        for _, p := range a {
            for {
                if i % p == 0 {
                    m[p] = true
                    i /= p
                } else {
                    break 
                }
            }
            
        }
        if i > 1 {
            m[i] = true
        }
    }
    return len(m)
}