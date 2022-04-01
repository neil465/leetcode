func expand(s string) []string {
    res := []string{}
    
    var g func(string,string)
    g = func(s, word string) {
        if len(s) == 0 {
            res = append(res, word)
            return
        }else{
            if string(s[0]) == "{" {
                index := 0
                
                for index = range s { 
                    if string(s[index]) == "}" {
                        break 
                    } 
                }
                
                for _,i := range strings.Split(s[1:index], ",") {
                    g(s[index + 1:], word + i)
                }
            }else{
                g(s[1:],word + string(s[0]))
            }
        }
    }


    g(s,"")
    sort.Strings(res)
    return res
}

