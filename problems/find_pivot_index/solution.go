func pivotIndex(nums []int) int {
    p := []int{}
    a := 0 
    for _, i := range nums {
        a += i
        p = append(p, a)
    }
    if p[len(nums)-1] - nums[0] == 0 {
        return 0
    }

   
    for i := 1; i < len(nums)-1 ; i++ {
        if p[i-1] == p[len(nums)-1] - p[i] {
            return i
        }
    }
     if len(nums) >= 2 && p[len(nums)-2] == 0 {
        return len(nums)-1
    }
    return -1
}