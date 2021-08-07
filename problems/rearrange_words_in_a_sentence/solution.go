func arrangeWords(text string) string {
    a := strings.Fields(text)
    sort.SliceStable(a,func(i,j int)bool{return len(a[i]) < len(a[j])})
    k := strings.Join(a, " ")
    k=strings.ToLower(k)
    pop := strings.ToUpper(string(k[0]) )
    k = k[1:]
    k = pop + k 
    return k
    
}