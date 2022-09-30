func capitalizeTitle(title string) string {
    res := ""
    
    for i := 0 ; i < len(title) ; i++{
        if (i == 0 || title[i - 1] == ' ' ) && i + 2 < len(title) && 
        (title[i + 1] != ' ' && title[i + 2] != ' '){ 
            res += strings.ToUpper(string(title[i ]))
        } else if title[i] == ' ' {
            res += " "
        } else {       
            res += strings.ToLower(string(title[i]))
        } 
    }
    return res
    
}