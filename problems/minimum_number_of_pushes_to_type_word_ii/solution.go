
import (
	"fmt"
	"sort"
)
func minimumPushes(word string) int {
    a := [26]int{}
    freq := map[rune]int{}

    for _, i := range word {
        freq[i] ++
    }
    arr := []rune(word)
    sort.Slice(arr, func(i, j int) bool {
        return freq[arr[i]] > freq[arr[j]]
    })

    word = string(arr)
    cur := 1
    v := 0
    for i := range word {
        if a[word[i] - 'a'] > 0 {
            v += a[word[i] - 'a']
        } else {
            depth := cur / 8
            if cur % 8 > 0 {
                depth ++
            }

            v += depth

            a[word[i] - 'a'] = depth
            cur ++
        }
    }

    return v
    
}