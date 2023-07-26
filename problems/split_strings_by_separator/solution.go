func splitWordsBySeparator(words []string, separator byte) []string {
    a := []string{}
    for _, i := range words {
        k := strings.Split(i, string(separator))
        for i := range k {
            if len(k[i]) > 0 {
                a = append(a, k[i])
            }
        }
        
    }
    return a
}