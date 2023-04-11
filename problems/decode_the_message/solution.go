func decodeMessage(key string, message string) string {
    p := "a"
    m := map[string]string{}
    for _, i := range key {
        if i != ' ' &&  m[string(i)] == ""{

            m[string(i)] = p
            p = string(int(p[0])+ 1)
        }
    }

    res := ""
    for _, i := range message {
        if i != ' ' {
            res += m[string(i)]
        } else {
            res += " "
        }
    }
    return res
}