func countBits(num int) []int {
    result := []int{}
    for i:=0;i<=num;i++{
        x := uint32(i)
        res := 0
        for x >0 {
            if x & uint32(1) == 1{
                res ++
            }
            x>>=1
        }
        result = append(result,res)
    }
    return result
}