func hammingDistance(x int, y int) int {
    res := 0
    ux := uint32(x)
    uy := uint32(y)
    for ux > 0 || uy > 0{
        if ux & uint32(1) != uy & uint32(1){
            res++
        }
        ux>>=1
        uy>>=1
    }
    return res
}