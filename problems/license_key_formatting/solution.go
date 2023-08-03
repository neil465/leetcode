func licenseKeyFormatting(s string, k int) string {
    s = strings.ReplaceAll(s, "-", "")

    res := strings.ToUpper(s[:len(s) % k] )
    s = s[len(s) % k:]

    for i := 0 ; i < len(s); i++ {
        if i % k == 0 {
            res += "-"
        }
        res += strings.ToUpper(string(s[i]))
    }
    return strings.Trim(res, "-")
   
    


    
}