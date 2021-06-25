func removeDuplicates(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return len(nums)
    }
    count:= 1
    for i:= 1 ; i < len(nums); i++{
        fmt.Println(i)
        if nums[i-1] != nums[i] {
            nums[count] = nums[i]
            count++
        }
    }
    return count

}
