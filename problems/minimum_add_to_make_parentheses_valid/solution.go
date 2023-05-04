func minAddToMakeValid(s string) int {
    need := 0
    res := 0
    for _, i := range s {
        if i == '(' {
            need ++
        } else {
            need--
        }
        if need < 0 {
            res ++
            need = 0
        } 
    }
    if need > 0 {
        return need + res
    }
    return  res
}