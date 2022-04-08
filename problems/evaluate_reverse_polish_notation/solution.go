func evalRPN(tokens []string) int {
    s := []int{}
    for _, token := range tokens {
        if token != "+" && token != "-" && token != "*" && token != "/" {
            intToken, _ := strconv.Atoi(token)
            s = append(s, intToken)
        }else{
            a , b := s[len(s) - 2],s[len(s) - 1]
            s = s[:len(s) - 2]
            if token == "+" {
                s = append(s, a + b)
            }else if token == "/"{
                s = append(s, a / b)
            }else if token == "*"{
                s = append(s, a * b)
            }else if token == "-"{
                s = append(s, int(float64(a) - float64(b)))
            }
        }
    }
    return s[0]
}