

func smallestString(s string) string {
    ac := []rune(s)
    start := -1
    for i := range s {
        if s[i] != 'a' {
            start = i
            break
        }

    }

    if start == -1 {
        return s[:len(s) - 1] + "z"
    }
    

    for start < len(s) && s[start] != 'a'{
        ac[start] --
        start ++
    }
    return string(ac)


}