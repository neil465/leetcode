func rearrangeArray(nums []int) []int {
    v1, v2 := []int{}, []int{}

    for _, i := range nums {
        if i < 0 {
            v1 = append(v1, i)
        } else {
            v2 = append(v2, i)
        }
    }

    cur := []int{}

    for i := 0 ; i < len(v1) ; i ++ {
        cur = append(cur, v2[i], v1[i])
    }
    return cur
}