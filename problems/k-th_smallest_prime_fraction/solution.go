
import "sort"
func kthSmallestPrimeFraction(arr []int, k int) []int {
    v := [][]int{}

    for i := range arr {
        for j := i + 1 ; j < len(arr); j++ {
            v = append(v, []int{arr[i], arr[j]})
        }
    }
    sort.Slice(v, func(i, j int) bool {
        return float64(v[i][0])/ float64(v[i][1]) < float64(v[j][0])/ float64(v[j][1])
    })


    return []int{v[k - 1][0], v[k - 1][1]}
}