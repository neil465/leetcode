func minStartValue(nums []int) int {
    sum := 1
    min := 100000000
    for _,j := range nums{
        sum += j
        if sum <= min{
            min = sum
        }
    }
    if min > 0 {
        return 1
    }
    return (min*-1)+2
}