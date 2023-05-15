func doesValidArrayExist(derived []int) bool {
    a := 0
    for _, i := range derived {
        a += i
    }
    return a % 2 == 0
}