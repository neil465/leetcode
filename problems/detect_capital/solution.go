func detectCapitalUse(word string) bool {
    uppercase := false
    lowercase := false
    for i,letter := range word{
        if unicode.IsUpper(letter) && i != 0{
            uppercase = true
            if lowercase{
                return false
            }
        }
        if unicode.IsLower(letter) {
            lowercase = true
            if uppercase{
                return false
            }
        }
    }
    return true
}