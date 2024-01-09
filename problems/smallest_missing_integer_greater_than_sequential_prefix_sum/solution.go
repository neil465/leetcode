func missingInteger(nums []int) int {
    m := map[int]bool{}

    for _, i := range nums {
        m[i]= true
    }

    s := 0 

    for i := 0 ; i < len(nums); i++ {
        if s != 0 && nums[i] -1 != nums[i - 1] {
            break
        }
        s += nums[i]
    }

    for i := s; i < math.MaxInt32 ; i++ {
        if !m[i]  {
            return i
        }
    }
    return -1
}