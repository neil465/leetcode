func numberGame(nums []int) []int {
    sort.Ints(nums)
    
    arr := []int{}
    
    for i := 0 ; i < len(nums); i += 2 {
        arr = append(arr, nums[i + 1])
        arr =append(arr, nums[i ])
    }
    return arr
}