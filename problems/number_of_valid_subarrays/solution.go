func validSubarrays(nums []int) int {
    sum := 1
    for i := 0 ; i < len(nums)-1 ; i++{
        for j := i ; j <len(nums) ;j++{
            if nums[j] < nums[i]{
                break
            }
            sum++
        }
    }
    return sum
}