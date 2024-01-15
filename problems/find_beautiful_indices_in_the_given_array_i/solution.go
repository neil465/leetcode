func beautifulIndices(s string, a string, b string, k int) []int {
    aind, bind := []int{}, []int{}

    for i := range s {
        if i + len(a) <= len(s)  && s[i:i + len(a)] == a {

            aind = append(aind, i)
        }
        if i + len(b) <= len(s) && s[i:i + len(b)] == b {
            bind = append(bind, i)
        }
    }
    fmt.Println(aind, bind)
    karr := []int{}

    for _, i := range aind {
        for _, j := range bind {
            if abs(i - j) <= k {
                karr = append(karr, i)
                break
            }
        }
    }
    return karr
}

func abs(i int)int {
    if i < 0 {
        return -i
    }
    return i
}