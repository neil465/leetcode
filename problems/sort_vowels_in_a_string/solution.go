func sortVowels(s string) string {
    a := []rune{}
    for _, i := range s {
        if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U' {
            a = append(a, i)
        }
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i] < a[j]
    })
    ind := 0
    k := []rune(s)
    for index := range s {
        i := s[index]
        if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U' {
            k[index] = a[ind]
            ind ++
        }
    }    
    return string(k)
}