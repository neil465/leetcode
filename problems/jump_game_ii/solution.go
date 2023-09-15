func jump(nums []int) int {
    cur := map[int]bool{0 : true}
    depth := 0
    for depth = 0 ; depth == depth ; depth ++ {

        t := map[int]bool{}

        for i := range cur {
            if i == len(nums) - 1{

                return depth
            }
            for j := i + 1; j < int(math.Min(float64(len(nums)), float64(i + nums[i] + 1))); j++ {

                t[j] = true
            }
        }
        cur = t
    }
    return -1
}
