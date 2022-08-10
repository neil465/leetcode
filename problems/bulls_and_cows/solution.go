func getHint(secret string, guess string) string {
    bulls,cows:= 0,0
    m := [100000]int{}
    for i := range secret {
        a, _ := strconv.Atoi(string(secret[i]))
                             b, _ := strconv.Atoi(string(guess[i]))
        if a == b {
            bulls++
        } else  {
            if m[a] > 0{
                cows ++
            }
            
            if m[b] < 0 {
                cows ++
            }
        } 
        m[a]--
        m[b]++
    }
    return fmt.Sprint(bulls)+"A"+fmt.Sprint(cows)+"B"
}