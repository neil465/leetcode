var dp = map[string]bool{}

func pyramidTransition(bottom string, allowed []string) bool {
    dp = map[string]bool{}
    if len(bottom)== 1 {
        return true
    }
    allowedMap := map[string][]string{}

    for _, str := range allowed {
        allowedMap[str[:2]] = append(allowedMap[str[:2]], string(str[2]))
    }

    
    return ptr(bottom, allowedMap)
}
func ptr(bottom string , allowedMap map[string][]string) bool {
    if _, ok := dp[bottom]; ok {
        return dp[bottom]
    }
    if len(bottom)== 1 {
        return true
    }

    for i := 0; i < len(bottom) - 1 ; i ++ {
        if _, ok := allowedMap[string(bottom[i]) + string(bottom[i + 1])] ; ! ok {
            dp[bottom] = false
            return false
        }
    }

    for _, possible := range findBases(0, bottom, allowedMap, "") {
        if ptr(possible, allowedMap) {
            dp[bottom] = true
            return true
        }
    }
    dp[bottom] = false
    return false
}

func findBases(i int, bottom string, allowedMap map[string][]string, cur string) []string{
    if i == len(bottom) - 1{
        return []string{cur}
    }
    res := []string{}
    
    if _, ok := allowedMap[string(bottom[i]) + string(bottom[i + 1])] ; ! ok {
        return []string{}
    }
    for _, val := range allowedMap[string(bottom[i]) + string(bottom[i + 1])] {
        res = append(res, findBases(i + 1, bottom, allowedMap, cur + val)...)
    }

    return res
    
}