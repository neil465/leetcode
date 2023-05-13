func countSeniors(details []string) int {
    a := 0
    for _,i := range details {
        b , _ := strconv.Atoi(i[11:13])

        if b > 60 {
            a ++
        }
    }
    return a
}