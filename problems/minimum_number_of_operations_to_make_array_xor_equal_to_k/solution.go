func minOperations(nums []int, k int) int {
    v := nums[0]

    for i := 1; i < len(nums); i++ {
        v ^= nums[i]
    }
    add := 0
    for cur := 0; cur < 32; cur ++ {
        curval := 1 << cur
        a, b := false, false

        if v & curval > 0 {
            a = true
        }
        if k & curval > 0 {
            b = true
        }
        if a != b {
            add ++
        }
    }
    return add
}