func decodeCiphertext(encodedText string, rows int) string {
    if rows == 1 {
        return encodedText
    }
    res := ""
    for i := 0 ; i < (len(encodedText) / rows) ; i++ {
        temp := ""

        temp += string(encodedText[i])


        
        v := i + (len(encodedText) / rows) + 1

        for v < len(encodedText){
            // fmt.Println("h1", len(encodedText))

            temp += string(encodedText[v])


            v += (len(encodedText) / rows) + 1


            
            
        }
        
        res += temp
    }
    return strings.TrimRight(res, " ")
}