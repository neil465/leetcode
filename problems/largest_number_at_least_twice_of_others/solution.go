func dominantIndex(nums []int) int {
    g := 0 
    k := 0
    for j,i := range nums{
        if i > g{
            g = i
            k = j
        }
    }
    flag := true
    for _,i := range nums{
        if i!= g{
            if 2*i > g{
                flag = false
            }
        }
    }
    if flag{
        return k
    }
    return -1
}