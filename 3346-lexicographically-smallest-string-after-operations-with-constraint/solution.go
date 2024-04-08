func getSmallestString(s string, k int) string {
    best := "abcdefghijklmnopqrstuvwxyz"

    result := ""

    for i := range s {
        bind := 0
        for bind = 0; bind < len(best) ; bind ++ {
            if dist(int(s[i] - 'a'), int(best[bind] - 'a')) <= k {
                k -= dist(int(s[i] - 'a'), int(best[bind] - 'a'))
                if k >= 0 {
                    result += string(best[bind])
                } else {
                    result += string(s[i])
                }
                break
            }
        }

        fmt.Println(string(best[bind]), dist(int(s[i] - 'a'), int(best[bind] - 'a')))
        
        
    }

    return result
}

func abs(i int)int {
    if i < 0 {
        return -i
    }
    return i
}
func dist(a, b int) int {
    return min(25 - a + b + 1, abs(b - a))
}
