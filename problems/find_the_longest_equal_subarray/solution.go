func longestEqualSubarray(nums []int, k int) int {
    if len(nums) == 1 {
        return 1
    }
    m := map[int]int{nums[0] : 1}
    i, j := 0, 0
    max := 1
    for i < len(nums) && j < len(nums){

        if j - (i - 1) - max > k {
            fmt.Println(j - (i - 1),  max,  k )
            m[nums[i]] --
            i ++
            continue 
        }
        j ++
        if j >= len(nums){
            break
        }
        m[nums[j]] ++
        if m[nums[j]] > max {
            max = m[nums[j]]
        } 
    }
    
    return max
}