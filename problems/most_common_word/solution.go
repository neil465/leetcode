func mostCommonWord(paragraph string, banned []string) string {
    ban := map[string]bool{}
    count := map[string]int{}
    
    max := -1
    maxStr := ""
    words := strings.FieldsFunc(paragraph, func(r rune) bool {
        return !unicode.IsLetter(r)
    })
    for _, banword := range banned {
        ban[banword] = true
    } 
    for _, i := range words {
        w := strings.ToLower(i)
        

        if ban[w] {
            continue 
        }
        count[w] ++
        if count[w] > max {
            max = count[w]
            maxStr = w
        }
    }
    return maxStr
}