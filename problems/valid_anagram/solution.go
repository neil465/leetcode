
import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	bytearrays := strings.Split(s, "")
	sort.Strings(bytearrays)
	bytearrays2 := strings.Split(t, "")
	sort.Strings(bytearrays2)
	string1 := strings.Join(bytearrays,"")
	string2 := strings.Join(bytearrays2,"")
	if string1 == string2 {
		return true
	}
	return false
}
