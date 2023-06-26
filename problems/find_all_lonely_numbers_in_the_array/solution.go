func findLonely(nums []int) []int {
    a := map[int]int{}
    b := []int{}
    for _, i := range nums {
        a[i]++
    }
    for _, i := range nums {
        if a[i - 1] == 0 && a[i + 1]== 0 && a[i] == 1  {
            b = append(b, i)
        }
    }
    return b
}