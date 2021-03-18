import "strconv"

func freqAlphabets(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if i+2< len(s)&& string(s[i+2]) == "#" {
			number, _ := strconv.ParseInt(s[i:i+2], 10, 0)
			letter := string(97 + number - 1)
			res += letter
            i+=2
		} else {
			number, _ := strconv.ParseInt(string(s[i]), 10, 0)
			letter := string(97 + number - 1)
			res += letter
		}

	}
	return res
}
