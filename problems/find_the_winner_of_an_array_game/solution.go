func getWinner(arr []int, k int) int {
    if k >= len(arr) {
        sort.Ints(arr)
        return arr[len(arr) - 1]
    }
    m := map[int]int{}
    for {
        if arr[0] > arr[1] {
            m[arr[0]] ++ 
            if m[arr[0]] == k {
                return arr[0]
            }
            arr = append(arr, arr[1])
            arr =  append(arr[:1], arr[2:]...)
            
            
        }
        if arr[0] < arr[1] {
            m[arr[1]] ++ 
            if m[arr[1]] == k {
                return arr[1]
            }
            arr = append(arr, arr[0])
            arr = arr[1:]
        }
    }
    return 0
}