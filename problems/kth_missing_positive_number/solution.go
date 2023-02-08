func findKthPositive(arr []int, k int) int {
    if len(arr) == 1 {
        if arr[0] != 1 {
            return arr[0] - 1
        } 
        return arr[0] + 1
    }
    arr = append([]int{0}, arr...)

    for i := 0 ; i < len(arr) - 1 ; i++ {
        if arr[i] != arr[i + 1] - 1 {
            if arr[i + 1] - 1 - arr[i] >= k {
                return arr[i] + k
            } else {
                k -= arr[i + 1] - 1 - arr[i]
            }
        }
        
    }

    return arr[len(arr) - 1] + k
}