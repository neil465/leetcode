func findLucky(arr []int) int {
    max := -1
    m := make(map[int]int)
    for _,i := range arr{m[i]++}
    for _,i := range arr{if m[i] == i && max < i{max = i}}
    return max
}