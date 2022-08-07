func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    m := map[int]int{}
    for _, i := range items1 {
        m[i[0]] += i[1]
    }
    for _, i := range items2 {
        m[i[0]] += i[1]
    }
    a := [][]int{}
    for i,j := range m{
        a = append(a, []int{i,j})
    }
    sort.Slice(a, func(i,j int)bool{
        return a[i][0]< a[j][0]
    })
    return a
}