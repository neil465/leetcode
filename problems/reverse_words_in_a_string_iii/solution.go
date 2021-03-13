import "strings"

func reverseWords(s string) string {
	res := ""
	splits := strings.Split(s, " ")
	for _, i2 := range splits {
		sum := ""
		for _, i3 := range i2 {
			sum = string(i3) + sum
		}
		res += sum + " "
	}
	return strings.TrimRight(res," ")
}
