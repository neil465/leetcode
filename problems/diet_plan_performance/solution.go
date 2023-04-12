func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
    s := 0
    v := 0
    for i := 0 ; i < k; i ++ {
        s += calories[i]
    }
    
    for i := k; i < len(calories); i ++ {
        if s < lower {
            v --
        }else if s > upper {
            v ++
        }
        s -= calories[i - k]
        s += calories[i]
    }
    if s < lower {
        v --
    }else if s > upper {
        v ++
    }
    return v
}