func countAlternatingSubarrays(nums []int) int64 {
    l := 0
    v := 0
    for i := range nums {
        if i == 0 || nums[i] != nums[i - 1] {
            l ++
        } else {
            v += (l * (l + 1)) / 2
            l = 1
        }
    }
    v += (l * (l + 1)) / 2
    return int64(v)
}
