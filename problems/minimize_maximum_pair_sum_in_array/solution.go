func minPairSum(nums []int) int {
    sort.Ints(nums)
    i,j := 0,len(nums)-1
    max := 0
    for i < len(nums)/2{
        if (nums[i] + nums[j]) > max{
            max = nums[i] + nums[j]
        }
        i ++
        j--
    }
    return max
    
}