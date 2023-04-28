func percentageLetter(s string, letter byte) int {
    a := 0
    for _, i := range s {
        if byte(i) == letter {
            a++
        }
    }
    if a == 0 {
        return 0
    }
    return int(float64(100) * float64(a) / float64(len(s)))
}