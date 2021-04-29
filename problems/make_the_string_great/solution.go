import "strings"

func makeGood(s string) string {
	b := []rune(s)
	for i := 0; i < len(b); i++ {
		if i >= len(b) {
			break
		} else {
            if i+1<len(b) && string(b[i]) != string(b[i+1]) && strings.EqualFold(string(b[i]),string(b[i+1])) {
				b = append(b[:i], b[i+1:]...)
				b = append(b[:i], b[i+1:]...)
				i = -1
			}
            
	}
		}
        
	s = ""
	for _, r := range b {
		s+= string(r)
	}
	return s
}