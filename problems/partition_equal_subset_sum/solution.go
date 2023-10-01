var dp = map[[3]int] bool{}

func canPartition(nums []int) bool {
    msum := 0
    for _, i := range nums {
        msum += i
    }
    msum /= 2 

     dp = map[[3]int] bool{}
    return re(0,0,0,nums, msum)
}
func re(s1, s2, ind int, nums []int, msum int) bool {
    if s1 > msum || s2 > msum {
        return false
    }
    if _,ok := dp[[3]int{s1, s2, ind}] ; ok {
        return dp[[3]int{s1, s2, ind}] 
    }
    if ind == len(nums){

        return s1 == s2
    }

    k := re(s1 + nums[ind ], s2, ind + 1, nums, msum) || re(s1, s2 + nums[ind], ind + 1, nums, msum)
    dp[[3]int{s1, s2, ind}] = k
    return k
}