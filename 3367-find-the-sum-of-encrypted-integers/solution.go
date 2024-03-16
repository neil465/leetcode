func sumOfEncryptedInt(nums []int) int {
    sum := 0 

    for _, v := range nums {
        best := -1
        l := 0
        for v > 0 {
            l++
            if v %10 > best {
                best = v % 10
            }
            v /= 10
        }
        curMulti := 1

        for i := 0 ; i < l; i ++ {
            sum += best * curMulti
            curMulti *= 10
        }
    }
    return sum
}
