func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    
    kidpos := 0
    cookiepos := 0
    
    for kidpos < len(g) && cookiepos < len(s){
        if s[cookiepos] >= g[kidpos]{
            kidpos++
        }
        cookiepos++
    }
    
    return kidpos
    
}