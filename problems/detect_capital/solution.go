func detectCapitalUse(word string) bool {
    if len(word) <= 1{
        return true
    }
    k := 0
    flag := true
    if unicode.IsLower(rune(word[0])) {
        flag = false
    }
    if unicode.IsLower(rune(word[0])) == false && unicode.IsLower(rune(word[1])){
        k = 1
        flag = false
    }
    
    for i := k  ; i <len(word) ; i++{
        if flag == true && unicode.IsLower(rune(word[i])){
            return false
        }
        if flag == false && unicode.IsLower(rune(word[i])) == false{
            return false
        }
    }
    return true
}