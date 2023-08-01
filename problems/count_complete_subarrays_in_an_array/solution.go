func countCompleteSubarrays(nums []int) int {
    m := map[int]bool{}
    for _, i := range nums {
        m[i] = true
    }
    c := 0
    for i := range nums {
        c += re(i, nums, map[int]bool{nums[i] : true}, len(m))
    }
    return c
}
func re(i int, nums []int, k map[int]bool, l int) int {
    c := 0
    if i == len(nums) {
        return 0
    }
    k[nums[i]] = true

    if len(k) == l {

        c ++
    }
    
    
    c += re(i + 1, nums, k, l)
    return c
}