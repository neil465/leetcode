func numberOfDays(year int, month int) int {
    months :=  []int{1, 3, 5, 7, 8, 10, 12}
    for _, i := range months {
        if month == i {
            return 31
        }
    }
    if month != 2 {
        return 30
    }
    res := 28
    if year % 4== 0 && (year % 100 != 0 || year % 400 == 0) {
        res ++
    }
    return res
}