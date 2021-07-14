func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    res := []int{}
    for _,i := range queries{
        nums[i[1]] += i[0]
        sum := 0
        for _,j := range nums{
            if j%2 ==0{
                sum += j
            }
        }
        res = append(res,sum)
    }
    return res
}