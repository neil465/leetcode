func areNumbersAscending(s string) bool {
    min := 0
    k := strings.Fields(s)
    for _,i := range k{
        b, _ := strconv.Atoi(i)
        if b != 0 {
            fmt.Println(b)
            if b <= min{
                return false
            }
            min = b
        }
    }
    return true
}