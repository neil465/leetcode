func findLengthOfLCIS(nums []int) int {
    if len(nums) == 1{
        return len(nums)
    }
    c,m := 1,0
    for i := 0 ; i < len(nums)-1 ; i++{
        if i == len(nums)-2{
            if nums[i] < nums[i+1]{
                c++
            }
            if c > m{
                return c
            }
        }
        if nums[i] < nums[i+1]{
            c++
        }else{
            if m < c{
                m = c
            }

            c = 1
        }
    }
    
    return m
}