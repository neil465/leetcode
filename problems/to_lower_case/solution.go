import "unicode"

func toLowerCase(str string) string {
	res := ""

	for _, i2 := range str {
		if unicode.IsLetter(rune(i2)) && i2 <= 90 {
				res += string(i2 + 32)
                continue 
        }
        res+= string(i2)
	}
	return res
}
