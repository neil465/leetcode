
import "strconv"
func evalRPN(tokens []string) int {
    s := []int{}

    for i := range tokens {
        if tokens[i] == "+" {
            v := s[len(s) - 1] + s[len(s) - 2]
            s = s[:len(s) - 2]
            s = append(s, v)
        } else if  tokens[i] == "*" {
            v := s[len(s) - 1] * s[len(s) - 2]
            s = s[:len(s) - 2]
            s = append(s, v)
        } else if tokens[i] == "/" {
            v := s[len(s) - 2] / s[len(s) - 1]
            s = s[:len(s) - 2]
            s = append(s, v)
        } else if tokens[i] == "-" {
            v := s[len(s) - 2] - s[len(s) - 1]
            s = s[:len(s) - 2]
            s = append(s, v)
        } else {
            v, _ := strconv.Atoi(tokens[i])
            s = append(s, v)
        }
    }
    return s[0]
}