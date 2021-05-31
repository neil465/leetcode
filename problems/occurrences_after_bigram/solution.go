func findOcurrences(text string, first string, second string) []string {
    res := []string{}
    texts := strings.Split(text," ")
    for i:= 0 ; i< len(texts)  ; i++{
        if i+2 < len(texts) && texts[i] == first && texts[i+1] == second{
            res = append(res,texts[i+2])
        } 
    }
    return res
}