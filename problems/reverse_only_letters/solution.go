func reverseOnlyLetters(s string) string {
    if len(s) < 2{
        return s
    }
    a := make([]string,len(s))
    l,r := 0 , len(s)-1
    for l <= r{
        if unicode.IsLetter(rune(s[l])) && unicode.IsLetter(rune(s[r])){
            a[r] = string(s[l])
            a[l] = string(s[r])
            l++
            r--
        }
        if !unicode.IsLetter(rune(s[l])) && unicode.IsLetter(rune(s[r])){ 
            a[l] = string(s[l])
            l++    
        }
        if unicode.IsLetter(rune(s[l])) && !unicode.IsLetter(rune(s[r])){
            a[r] = string(s[r])
            r--
        }
        if !unicode.IsLetter(rune(s[l])) && !unicode.IsLetter(rune(s[r])){
            a[l] = string(s[l])
            a[r] = string(s[r])
            l++
            r--
        }
    }
    return strings.Join(a,"")
}