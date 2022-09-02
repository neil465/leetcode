func edgeScore(edges []int) int {
    max, index, m := 0,0, map[int]int{}
    for i,j := range edges {
        m[j] += i
        if m[j] > max {
            max = m[j]
            index = j
        }
        if m[j] == max && j < index{
            index = j
        }
    }
    fmt.Println(m)
    return index
}