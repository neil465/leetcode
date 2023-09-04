func canBeEqual(s1 string, s2 string) bool {
    m := map[string]bool{}
    a := []rune(s1)
    a[0], a[2] = a[2], a[0]
    m[string(a)] = true
    a[1], a[3] = a[3], a[1]
    m[string(a)] = true 
    a[0], a[1], a[2], a[3] = a[2], a[3], a[0], a[1]
    a[1], a[3] = a[3], a[1]
    m[string(a)] = true 
    a = []rune(s2)
    a[0], a[2] = a[2], a[0]
    if m[string(a)] {
        return true
    }
    a[1], a[3] = a[3], a[1]
    if m[string(a)] {
        return true
    }
    a[0], a[1], a[2], a[3] = a[2], a[3], a[0], a[1]
    a[1], a[3] = a[3], a[1]
    if m[string(a)] {
        return true
    }
    return false
}
 