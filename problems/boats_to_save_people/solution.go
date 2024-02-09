func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    i, j := 0, len(people) -1
    b := 0
    for i <= j {
        if people[j] + people[i] > limit {
            i --
        }
        j --
        i ++
        b++
    }
    return b
}