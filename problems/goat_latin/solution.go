func toGoatLatin(sentence string) string {
    m := map[string]bool{"a":true, "e":true,"i":true,"o":true,"u":true,"A":true, "E":true,"I":true,"O":true,"U":true}

    arr := strings.Fields(sentence)

    res := ""

    for index, i := range arr {
        if m[string(i[0])] {
            res += i + "ma" + strings.Repeat("a", index + 1) + " "
        } else {
            res += i[1:] + string(i[0])+ "ma"+ strings.Repeat("a", index + 1) + " "
        }
    }
    return res[:len(res) - 1]
}