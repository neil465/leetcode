func maxDistance(colors []int) int {
    i,j := 0,len(colors)-1
    for colors[len(colors)-1] == colors[i]{
        i++
    }
    for colors[j] == colors[0]{
        j--
    }
    return max(len(colors)-i-1,j)
}
func max(i,j int)int{
    if i > j {return i}
    return j
}