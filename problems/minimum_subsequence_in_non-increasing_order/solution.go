func minSubsequence(nums []int) []int {
    sort.Ints(nums)
    if len(nums) == 1{
        fmt.Println("h")
        return nums
    }
    for i:= len(nums)-1 ; i >=0 ; i--{
        sum := 0
        sum2 := 0
        for _,j := range nums[i:]{
            sum += j
        }
        for _,j := range nums[:i]{
            sum2 += j
        }
        
        if sum > sum2{
            n:=nums[i:] 
            sort.Sort(sort.Reverse(sort.IntSlice(n)))
            return n
        }
    }
    return []int{}
}