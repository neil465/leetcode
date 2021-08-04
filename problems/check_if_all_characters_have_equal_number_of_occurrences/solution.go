func areOccurrencesEqual(s string) bool {
    m,a := make(map[rune]int) , -120012324
    for _,i := range s{
        m[i]++
    }
    for _,i := range m{
        if i != a && a == -120012324{
            a = i
        }else if i != a && a != -120012324{
            return false
        }
    }
    return true
}