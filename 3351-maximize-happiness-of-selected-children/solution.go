func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Slice(happiness, func(i, j int) bool {
        return happiness[i] > happiness[j]
    })

    v := 0

    for i := 0 ;i < k; i ++ {
        v += max(happiness[i] - i, 0)
    }

    return int64(v)
}
