func furthestDistanceFromOrigin(moves string) int {
    lc, rc, dcount := 0, 0 , 0
    
    for _, i := range moves {
        if i == 'L' {
            lc ++
        } else if i == 'R' {
            rc ++
        } else {
            dcount ++
        }
    }
    
    if lc > rc {
        return lc - rc + dcount
    } 
    return rc - lc + dcount
}