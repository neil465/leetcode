func countDistinctIntegers(nums []int) int {
    m := map[int]bool{}
    for _,i := range nums {
        m[i] = true
        revnum :=  0
        for i > 0 {
            revnum, i = revnum * 10 + i%10 , i/10
        }
        m[revnum] = true
        
    }
    return len(m)
}