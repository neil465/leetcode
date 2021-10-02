func rotate(nums []int, k int)  {
    
    v := k
    if v > len(nums)-1{
        v = v%len(nums)
    }
    l,r := 0,len(nums)-1
    for l<=r{
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
    l,r = 0,v-1
    
    for l<=r{
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
    l,r = v,len(nums)-1

    for l<=r{
        nums[l],nums[r] = nums[r],nums[l]
        l++
        r--
    }
    
    
       
}

