func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    a := 0
    for i := range seats {
        a += abs(seats[i]-students[i])
    }
    return a
}
func abs (i int) int {
    if i < 0 {
        return -i
    }
    return i

}