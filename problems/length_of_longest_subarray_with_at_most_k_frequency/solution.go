func maxSubarrayLength(nums []int, k int) int {
    maxVal:= -1
    m := map[int]int{}
    i := 0
    best := 0 
    for j := range nums {
        m[nums[j]] ++

        if m[nums[j]] > m[maxVal] {
            maxVal = nums[j]
        }

        for m[maxVal] > k {
            m[nums[i]] --
            i++
        }
        if j - i > best {
            best = j - i
        }
    }
    return best + 1
}

