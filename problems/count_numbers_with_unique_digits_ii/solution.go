
import "fmt"
func numberCount(a int, b int) int {
    count := b - a + 1
    for i := a; i <= b; i ++ {
        v := i
        m := map[int]bool{}
        for v > 0 {

            if m[v % 10] {
                break
            }
            m[v % 10] = true
            
            v /= 10
        }
        if v > 0 {
            count --
        }
    }   
    return count
}
