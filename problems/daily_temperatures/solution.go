type key struct {
    k int
    v int
}


func dailyTemperatures(temperatures []int) []int {
    res := make([]int,len(temperatures))
    ms := []key{}
    for i := len(temperatures)-1 ; i >= 0; i--{
        for len(ms) > 0 && temperatures[i] >= ms[len(ms)-1].k{
            ms = ms[:len(ms)-1]
        }
        res[i] = 0
        if len(ms) > 0 {
            res[i] = ms[len(ms)-1].v - i
        }
        ms = append(ms,key{k:temperatures[i],v:i})
    }
    return res
}


