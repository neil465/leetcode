func wiggleSort(nums []int)  {
    sort.Ints(nums)
    for i := 2 ; i< len(nums) ; i++{
        if (i+1)%2 != 0{
            nums[i],nums[i-1] = nums[i-1],nums[i]
        } 
    }
}