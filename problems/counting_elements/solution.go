func countElements(arr []int) int {
    nums := map[int]bool{}

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
    sum := 0
    for _, i := range arr {
        if nums[i + 1] {
            sum ++
        }
        nums[i] = true
    }
    return sum

}
