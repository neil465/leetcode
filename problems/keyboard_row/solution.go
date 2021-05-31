func findWords(words []string) []string {
    res := []string{}
    
    v1 := "qwertyuiop"
    v2 := "asdfghjkl"
    v3 := "zxcvbnm"
    m := make(map[string]int)
    for _,i:= range v1{
        m[string(i)] = 1
    }
    for _,i:= range v2{
        m[string(i)] = 2
    }
    for _,i:= range v3{
        m[string(i)] = 3    
    }
    for _,j := range words{
        flag := true
        t := m[strings.ToLower(string(j[0]))]
        fmt.Println(m)
        for i := 1 ; i< len(j) ; i++{
            if m[strings.ToLower(string(j[i]))] != t{
                flag = false 
                break
            }
        }
        if flag{
            res = append(res,j)
        }
    }
    return res
}