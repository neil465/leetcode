func kthSmallest(matrix [][]int, k int) int {
    arr := []int{} 
    for _,i := range matrix{
        for _,i2 := range i{
            arr = append(arr,i2)
        }
    }
    sort.Ints(arr)
    return arr[k-1]
}