func stringShift(s string, shift [][]int) string {
    pos := 0
    for _,i := range shift{
        if i[0] == 1{
            pos+= i[1]
        }else{
            pos -= i[1]
        }
    }
    
    pos = pos % len(s)
    var offset int
    if offset = 0-pos; pos > 0{
        offset = len(s)-pos
    }
    
    return s[offset:]+ s[:offset]
}