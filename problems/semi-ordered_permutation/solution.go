func semiOrderedPermutation(nums []int) int {
    add := 0
    oi := 0
    for i := range nums {
        if nums[i] ==1 {
            oi = i 
            break
        }
    }
    for i := oi - 1 ; i >= 0 ; i -- {
        nums[i + 1], nums[i] = nums[i], nums[i + 1]
        add ++ 
    }
    for i := range nums {
        if nums[i] == len(nums) {
           oi = i 
        }
    }
    for i := oi+1 ; i < len(nums) ; i ++ {
        nums[i - 1], nums[i] = nums[i], nums[i - 1]
        add ++ 
    }
    return add
    
}