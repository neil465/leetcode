func maxAscendingSum(nums []int) int {
    greatest := 0
    for j,i := range nums{
        sum := i
        now := i
        if j+1 < len(nums){
            for k := j+1 ; k < len(nums); k++{
                if nums[k] <=now {
                    break
                }
                sum += nums[k]
                now = nums[k]
                fmt.Println(sum)
            }
        }
        if sum > greatest{
            fmt.Println("_________")
            greatest =sum
        }
        
    }
    return greatest
}