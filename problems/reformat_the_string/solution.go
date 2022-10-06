func reformat(s string) string {
    letters, numbers, result  := "","", ""

    for i := 0; i < len(s) ; i++ {
        if unicode.IsLetter(rune(s[i])){
            letters += string(s[i])
        } else {
            numbers += string(s[i])
        }
    }
    if len(letters) - len(numbers) < -1 || len(letters) - len(numbers) > 1 {
        return ""
    }
    add := len(letters) > len(numbers)
    for i := 0;  i < len(s) ; i++{
        if add {
            result, letters = result + string(letters[0]), letters[1:]
        } else {
            result, numbers = result + string(numbers[0]), numbers[1:]
        }
        fmt.Println(result)
        add = !add
    }
    return result
}