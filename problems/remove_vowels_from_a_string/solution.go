func removeVowels(s string) string {
    s=strings.ReplaceAll(s,"a","")
    s=strings.ReplaceAll(s,"e","")
    s=strings.ReplaceAll(s,"i","")
    s=strings.ReplaceAll(s,"o","")
    s=strings.ReplaceAll(s,"u","")
    return s
}