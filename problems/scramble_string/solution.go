var dp = map[[2]string]bool{}

func isScramble(s1 string, s2 string) bool {
    dp = map[[2]string]bool{}
    return re(s1, s2)
}
func re(s1, s2 string) bool {
    if _, ok := dp[[2]string{s1,s2}]; ok {

        return dp[[2]string{s1,s2}]
    }
    if len(s1) == 1 {
        return s1 == s2
    }
    if s1 == s2 {
        return true
    }
    v := false
    for i := 1; i < len(s1) ; i++ {
        if (re(s1[:i], s2[:i]) && re(s1[i:], s2[i:])) || (re(s1[i:], s2[:len(s2) - i]) && re(s1[:i], s2[len(s2) - i :])) {
            v = true
            break
        }
    }
    dp[[2]string{s1,s2}] = v
    return v
}
