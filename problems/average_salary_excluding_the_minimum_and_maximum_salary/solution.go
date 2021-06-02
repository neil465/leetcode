func average(salary []int) float64 {
    sort.Ints(salary)
    miniumum := salary[0]
    maximum := salary[len(salary)-1]
    res := 0
    for _, i := range salary{
        res += i
    }
    num := float64(res-maximum-miniumum)
    div := float64(len(salary)-2)
    return float64(num/div) 
    
}