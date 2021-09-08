func findMiddleIndex(nums []int) int {
    if len(nums) == 0{return -1}
    if len(nums) < 2{return 0}
    sum := 0
    prefix := []int{}
    
    for _,i := range nums{
        sum += i
        prefix = append(prefix,sum)
    }
    for j,i := range prefix{
        if j == 0{
            if prefix[len(prefix)-1]-prefix[0] == 0{
                return j
            }
            continue
        }
        if j == len(prefix)-1{
            if prefix[len(prefix)-2] == 0{
                return j
            }
            continue
        }
        if prefix[len(prefix)-1] - i == prefix[j-1]{return j}
    }
    return -1
}