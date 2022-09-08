func isSelfCrossing(distance []int) bool {
    b,c,d,e,f := 0,0,0,0,0
    for i := 0 ; i < len(distance) ; i++{
        a := distance[i]
        if (d >= b && b > 0) && (a >= c || (a >=  c- e && c-e >= 0 && f >= d-b)) {
            return true
        }
        b,c,d,e,f = a,b,c,d,e
    }
    return false
}