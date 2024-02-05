import "strings"
func minimumTimeToInitialState(word string, k int) int {
    for i := k ; i < len(word) ; i += k {
        if strings.HasPrefix(word, word[i:]) {
            return i / k
        }
    }
    return int(math.Ceil(float64(len(word))/ float64(k)))
}