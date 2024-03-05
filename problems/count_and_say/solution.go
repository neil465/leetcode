
import "fmt"
func countAndSay(n int) string {
    str := "1"

    for i := 1 ; i < n ; i ++ {
        
        cnt := 0
        prev := rune(str[0])
        next := ""
        for _, v := range str {
            if v != prev {
                next += fmt.Sprint(cnt) + string(prev)
                cnt = 1
                prev = v
            } else {
                cnt ++
            }
        }

        if cnt > 0 {
            next += fmt.Sprint(cnt) + string(prev)
        }
        str = next

    }

    return str
}