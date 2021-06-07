func distributeCandies(candyType []int) int {
    m := make(map[int]int)
    for _,i := range candyType{
        m[i]++
    }
    if len(m) >= len(candyType)/2{
        return len(candyType)/2
    }
    return len(m)
}