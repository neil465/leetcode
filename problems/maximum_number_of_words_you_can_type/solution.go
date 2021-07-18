func canBeTypedWords(text string, brokenLetters string) int {
    m := make(map[rune]bool)
    res := 0
    
    for _,i := range brokenLetters{
        m[i] = true
    }
    
    for _,i := range strings.Split(text," "){
        flag := true
        for _,j := range i{
            if m[j] == true{
                flag = false
                break
            }
        }
        if flag {
            res++
        }
    }
    return res
}