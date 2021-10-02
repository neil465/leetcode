func sortedSquares(nums []int) []int {
    if len(nums) == 1{
        nums[0] = nums[0]*nums[0]
        return nums
    }
    k := make([]bool , len(nums))
    sort.Slice(nums,func(i,j int)bool{
        if k[i] != true{
            nums[i] = nums[i]*nums[i]            
            k[i] = true
        }
        if k[j] != true{
            nums[j] = nums[j]*nums[j]
            k[j] = true
        }
        
        return nums[i]< nums[j]
    })
    return nums
}   