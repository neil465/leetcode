func isMajorityElement(nums []int, target int) bool {
    flag := false
    for _,i := range nums{
        if i == target{
            flag = true
            break
        }
    }
    if flag == false {
        return false
    }
    if len(nums) == 1{
        return nums[0] == target
    }
    l := len(nums)
    for len(nums) > 0{
        fmt.Println(nums)
        if nums[0] == target && nums[len(nums)-1] == target{
            return len(nums) > l/2
        }else if nums[0] != target && nums[len(nums)-1] == target{
            nums = nums[1:]
        }else if nums[0] == target && nums[len(nums)-1] != target{
            nums = nums[:len(nums)-1]
        }else{
            nums = nums[1:len(nums)-1]
        }
    }
    return false
}