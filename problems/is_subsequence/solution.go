func isSubsequence(s string, t string) bool {
    pointer := 0
    if s == ""{
        return true
    }
    if t == "" {
        return false
    }
    for i := 0 ; i < len(t) ; i++{
        if s[pointer] == t[i]{
            pointer++
            if pointer >= len(s){
                return true
            }
        }
    }
    return pointer == len(s)
}