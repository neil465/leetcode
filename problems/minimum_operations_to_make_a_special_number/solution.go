import (
    "fmt"
    "strconv"
)

func minimumOperations(num string) int {
    k := false
    v := len(num)
    for i := range num {
        if num[i] == '0' {
            k = true
        }
        for j := i + 1 ; j < len(num); j++ {
            a, _ := strconv.Atoi(string(num[i]))
            b, _ := strconv.Atoi(string(num[j]))
            if a == 0 {
                k = true
            }
            if (a * 10 + b ) % 25 == 0 {
                if len(num) - i - 2 <= v {
                    v  =len(num) - i - 2
                    fmt.Println(i, j)
                }
            }
        }
    }
    if k && v == len(num) {
        return len(num) - 1
    }
    return v



}
