func minFlips(target string) int {
    target = strings.TrimLeft(target, "0")
    if len(target) == 0 {
        return 0
    }
    g, n := 0, byte('-')
    for i := 0 ; i < len(target);i ++ {
        if target[i] != n {
            n = target[i]
            g ++
        }
    }
 
    if g ==0 {
        return 0
    }
    return g 
}

/*
10110
01011
*/