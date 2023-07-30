func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    k := 0 
    for _, i := range hours {
        if i >= target {
            k ++
        }
    }
    return k
}