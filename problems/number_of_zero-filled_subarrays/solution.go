func zeroFilledSubarray(nums []int) int64 {
    sum := 0 
    l := 0
    for i := 0 ; i < len(nums) ; i++ {
        fmt.Println(l)
        if nums[i] == 0 {
            l++
        }else {
            sum += (l*(l+1))/2
            l= 0
        }
        
        
    }
    if l != 0 {
        sum += (l*(l+1))/2
    }
    return int64(sum)
}