func interchangeableRectangles(rectangles [][]int) int64 {
    count := int64(0)
    m := make(map[float64]int)
    for _,i := range rectangles{
        temp := float64(i[0])/float64(i[1])
        count += int64(m[temp])
        m[temp]++
    }

    return count
}