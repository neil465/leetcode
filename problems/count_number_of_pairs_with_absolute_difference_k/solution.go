func countKDifference(nums []int, k int) int {
    count := 0
    for i := range nums{
        for j := i+1 ; j<len(nums) ; j++{
            if abs(nums[i]-nums[j]) == k{
                count++  
            }
        }
    }
    return count
}
func abs (i int) int{
    if i < 0{return i * -1}
    return i    
}