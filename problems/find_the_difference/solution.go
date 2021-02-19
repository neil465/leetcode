import (
	"sort"
	"strings"
)

func findTheDifference(s string, t string) byte {
	sarr,tarr := strings.Split(s, ""),strings.Split(t, "")
	sort.Strings(sarr)
	sort.Strings(tarr)
	s,t = strings.Join(sarr,""),strings.Join(tarr,"")
	

    if len(s) > len(t){
        s,t = t,s
    }
	if len(s) < 1 {
		return t[0]
	}
	for i, i2 := range s {
		if string(i2) != string(t[i]) {
			return t[i]
		}
	}
	if s[len(s)-1] != t[len(t)-1] {
		return t[len(t)-1]
	}
	return t[len(t)-1]
}   
