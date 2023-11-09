func countHomogenous(s string) int {
    cnt, res := 0,0  
    for i := 0 ; i < len(s) + 1; i ++ {
        if i == len(s) || s[i] != s[i - cnt] { res, cnt = (cnt * (cnt + 1)) / 2 + res, 0 }
        cnt ++
    }
    return res % 1000000007
}