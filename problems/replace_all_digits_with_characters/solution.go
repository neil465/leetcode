import (
	"strconv"
	
	"unicode"
)

func replaceDigits(s string) string {
	b := []rune(s)
	for i := 0; i < len(b); i++ {
		if !unicode.IsLetter(b[i]) {
			j, _ := strconv.Atoi(string(b[i]))
			b[i] = rune(s[i-1] + uint8(j))
		}
	}
	sum := ""
	for _, r := range b {
		sum+=string(r)
	}
	return sum
}