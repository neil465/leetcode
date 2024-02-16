func findLeastNumOfUniqueInts(arr []int, k int) int {
    a := []int{}
    m := map[int]int{}

    for _, i := range arr {
        m[i] ++
    }

    for _, c := range m {
        a = append(a, c)
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i] < a[j]
    })
    for i := 0 ; i < len(a); i ++ {
        if a[i] > k {
            return len(a) - i
        }
        k -= a[i]
    }

    return 0
}