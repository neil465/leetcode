func buildArray(nums []int) []int {
    arr := []int{}
    for i,_ := range nums{
        arr = append(arr,nums[nums[i]])
    }
    return arr
}