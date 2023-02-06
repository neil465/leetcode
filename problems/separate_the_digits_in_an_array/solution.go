func separateDigits(nums []int) []int {
    res := []int{}

    for _, i := range nums {
        temp := []int{}
        for i > 0 {
            temp = append([]int{i%10}, temp...)
            i /= 10
        }
        res = append(res, temp...)

    }
    return res
}