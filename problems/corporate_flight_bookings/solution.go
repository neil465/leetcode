func corpFlightBookings(bookings [][]int, n int) []int {
    res := make([]int, n)
    for _, i := range bookings {
        for j := i[0]; j <= i[1]; j ++ {
            res[j - 1] += i[2]
        }
    }
    return res
}