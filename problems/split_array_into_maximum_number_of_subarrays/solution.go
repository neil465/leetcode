var dp = map[[3]int][]int{}

func maxSubarrays(nums []int) int {
    v := nums[0]
    for i := 1; i < len(nums); i++ {
        v = v & nums[i]
    }
    if v != 0 {
        return 1
    }
    cursum, k := nums[0], 0
    for i := 1; i < len(nums); i++ {
        if cursum == 0 {

            k ++ 
            cursum = nums[i]
        }  
        cursum = cursum & nums[i]
    }
    if cursum == 0 {
        k++
    }
    return k
}

