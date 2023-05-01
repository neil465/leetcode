func countPoints(points [][]int, queries [][]int) []int {
    res := []int{}
    for _, c := range queries {
        sum := 0
        for _, p := range points {
            if distance(p[0], p[1], c[0], c[1]) <= float64(c[2]) {
                sum++
            }
        }
        res = append(res, sum)
    }
    return res
}
func distance(x, y, x2, y2 int) float64 {
    return math.Sqrt(float64((x2 - x) * (x2 - x) + (y2 - y) * (y2 - y)))
}