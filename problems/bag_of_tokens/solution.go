
import "fmt"
func bagOfTokensScore(tokens []int, power int) int {
    sort.Ints(tokens)
    
    

    s := 0    
    for power >= 0 && len(tokens) > 0{
        if power >= tokens[0] {
            s++
            power -= tokens[0]
            tokens = tokens[1:]
        } else {
            if len(tokens) == 1 || s <= 0{
                break
            }

            power += tokens[len(tokens) - 1]

            s--

            tokens = tokens[:len(tokens) - 1]
        }
    }
    return s
}