func numberOfPairs(nums []int) []int {
    if len(nums) == 1 {
        return []int{0,1}
    }
    pairs, left := 0,0
    sort.Ints(nums)
    for i := 0 ; i < len(nums) ; i++{
        
        if i != len(nums)-1 && nums[i] == nums[i+1]  {
            pairs ++
            i++
        }else {
            left++
        }
    }
    return []int{pairs,left}
    
}