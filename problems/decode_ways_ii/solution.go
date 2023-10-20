var sm = map[string]bool{"1" : true, "2" : true, "3" : true, "4" : true, "5" : true, "6" : true, "7": true, "8" : true, "9" : true, "10" : true, "11" : true, "12" : true, "13" : true, "14" : true, "15" : true, "16" : true, "17": true, "18" : true, "19" : true, "20" : true, "21" : true, "22" : true, "23" : true, "24" : true, "25" : true, "26" : true}


func numDecodings(s string) int {
    return help(s, 0, map[int]int{}) % 1000000007
}
func help(s string, i int, dp map[int]int) int {
    if i == len(s) {
        return 1
    }
    if s[i] == '0' {
        return 0
    }
    if _, ok := dp[i]; ok {
        return dp[i] % 1000000007
    }
    
    v := 1 
    if s[i] == '*' {
        v = 9
    }
    v *=  help(s, i + 1, dp)

    if i + 1 < len(s) {
        if sm[s[i: i + 2]] {
            v += help(s, i + 2, dp)
        } else if s[i + 1] == '*' && s[i] == '*'{
            k := 15
            
            v += help(s, i + 2, dp) * k
        } else if s[i + 1] == '*' {
            k := 0
            if s[i] == '1' {
                k = 9
            } else if s[i] == '2' {
                k = 6
            }
            v += help(s, i + 2, dp) * k
        } else if s[i ] == '*' {
            k := 1
            
            if s[i + 1] <= '6' {
                k ++
            } 

            
            v += help(s, i + 2, dp) * k
        }
        
    }
    dp[i] = v
    return v % 1000000007
}