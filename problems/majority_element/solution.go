func majorityElement(nums []int) int {
    sum , num := 0,nums[0]
    for i := 0 ; i < len(nums) ; i++{
        if sum == 0 {
            num = nums[i]
        }
        if nums[i] == num{
            sum++
        }else{
            sum--
        }
    }
    return num
}