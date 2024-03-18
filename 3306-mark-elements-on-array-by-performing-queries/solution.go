func unmarkedSumArray(nums []int, queries [][]int) []int64 {
    type k struct {
        val, ind int
    }
    arr := []k{}

    cursum := 0

    for i := range nums {
        arr = append(arr, k{nums[i], i})
        cursum += nums[i]
    }

    sort.SliceStable(arr, func(i, j int) bool {
        return arr[i].val < arr[j].val
    })

    used := make([]bool, len(arr))

    indChange := make([]int, len(arr))

    for i := range arr {
        indChange[arr[i].ind] = i
    }

    res := []int64{}
    last := 0
    for _, q := range queries {
        if !used[indChange[q[0]]] {
            cursum -= arr[indChange[q[0]]].val
            used[indChange[q[0]]] = true
        }
        for i := last ; i < len(arr) && q[1] > 0; i ++ {
            if !used[i] {
                q[1] --
                used[i] = true
                cursum -= arr[i].val
                last = i + 1
            }
        }
        res = append(res, int64(cursum))
    }

    return res
}
