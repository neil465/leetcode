func maxSpending(values [][]int) int64 {
    inds := make([]int, len(values))
    val := 0
    for  i := range inds {
        inds[i] = len(values[i]) - 1
    }
    for i := 1 ; i < math.MaxInt32 ; i ++ {
        minind := math.MaxInt32
        minshop := math.MaxInt32
        for shop, ind := range inds {
            if ind >= 0 && (minind == math.MaxInt32 || values[shop][ind] < values[minshop][ minind]  ){
                minind = ind
                minshop = shop
            }
        }
        if minind == math.MaxInt32 {
            fmt.Println(inds)

            return int64(val)
        }
        // fmt.Println(inds)
        val += values[minshop][minind] * i
        inds[minshop] --
    }
    return int64(val)
}