func optimalDivision(nums []int) string {
    if len(nums) ==  1{
        return fmt.Sprint(nums[0])
    }
    if len(nums) ==2 {
        return fmt.Sprint(nums[0]) + "/" + fmt.Sprint(nums[1])
    }
    res := fmt.Sprint(nums[0])+"/"+"("
    for i := 1 ; i < len(nums) ; i++{
        res += fmt.Sprint(nums[i])+"/"
    }
    res = res[:len(res)-1]
    res += ")"
    return res
}