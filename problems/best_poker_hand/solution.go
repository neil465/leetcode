func bestHand(ranks []int, suits []byte) string {
    m,c := map[byte]int{},0
    for _,i := range suits {
        if m[i]+1 > c {
            c = m[i] + 1
        }    
        m[i]++
    }
    if c >= 5 {
        return "Flush"
    }
    nm,nc := map[int]int{},0
    for _,i := range ranks {
        if nm[i]+1 > nc {
            nc = nm[i] + 1
        }    
        nm[i]++
    }
    if nc >= 3 {
        return "Three of a Kind"
    }else if nc >= 2 {
        return "Pair"
    }
    if len(suits)>0 {
        return"High Card"
    }
    return ""
}