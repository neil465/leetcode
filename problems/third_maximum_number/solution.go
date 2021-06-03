func thirdMax(nums []int) int {
    sort.Ints(nums)
    fmt.Println(nums)
    
    for i := 0 ; i< len(nums) ; i++{
        if i!= 0 && nums[i] == nums[i-1]{
            nums=append(nums[:i], nums[i+1:]...)
            i--
        }
    }
    if len(nums) < 3{
        return nums[len(nums)-1]
    }
    return nums[len(nums)-3]
}