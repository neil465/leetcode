func countValidWords(sentence string) int {
    a := strings.Fields(sentence)
    v := len(a)
    fmt.Println(a)
    for i := range a {
        hc, pc := 0, 0
        for j := range a[i] {
            fmt.Println(hc, pc)
            if hc >= 2 || pc >= 2 {
                v --
                break
            }
            if a[i][j] == '-' {
                hc ++
                if ((j == 0 || j == len(a[i]) - 1) || !unicode.IsLetter(rune(a[i][j - 1])) || !unicode.IsLetter(rune(a[i][j + 1]))) {
                    v --

                    break
                }
                
            } else if unicode.IsDigit(rune(a[i][j])) {
                v -- 

                break
            } else if (string(a[i][j]) == "!" || string(a[i][j]) == "." || string(a[i][j]) == ",") {            
                pc ++
                if j != len(a[i]) - 1 {
                    v --

                    break
                }
                
            }
        }
    }
    return v
}