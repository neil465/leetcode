func largestUniqueNumber(nums []int) int {
    arr := make([]int,1001)
    for _,i := range nums{
        arr[i]++
    }
    for i := len(arr)-1 ; i > 0 ; i--{
        if arr[i] == 1{
            return i
        }
    }
    return -1
}