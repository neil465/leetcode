func repeatedCharacter(s string) byte {
    m := map[rune]bool{}
    for _,i := range s {
        if m[i]{
            return byte(i)
        }        
        m[i] = true
    }
    return byte('a')
}