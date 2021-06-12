func hammingWeight(num uint32) int {
    f := uint32(1)
    res := 0
    for num >0{
        if num & f == 1{
            res++
        }
        num = num>>1
    }
    return res
}