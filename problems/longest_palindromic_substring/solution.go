var dp = [1001][1001]string{}

func longestPalindrome(s string) string {
    dp = [1001][1001]string{}
    a := helper(s, 0, 0)
    if a == "@" {
        return ""
    }
    return a
}
func helper(s string, i, j int) string {
    if i > j ||j >= len(s) {
        return "@"
    }

    if dp[i][j] != "" {
        return dp[i][j]
    }
    max := "@"
    if isPalindrome(s, i, j) {

        max = s[i:j + 1]
    }
    one, two := helper(s, i + 1, j), helper(s, i, j + 1)
    if len(one) >= len(max) && one != "@"{
        max = one
    }
    if len(two) >= len(max) && two != "@" {
        max = two
    }
    dp[i][j] = max
    return max
}
func isPalindrome(st string, i, j int) bool {
    for s, e := i, j ; s < e ; s, e = s + 1, e - 1  {
        if st[s] != st[e] { return false }
    }
    return true
}