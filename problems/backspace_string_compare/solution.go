import "strings"

func backspaceCompare(s string, t string) bool {
	stack := []string{}
	for _, i2 := range s {
		if len(stack) > 0 && string(i2) == "#" {
			stack = append(stack[:len(stack)-1], stack[len(stack):]...)
		} else if string(i2) == "#"{
			continue
		}else {
			stack = append(stack, string(i2))
		}
	}
	stack2 := []string{}
	for _, i2 := range t {
		if len(stack2) > 0 && string(i2) == "#" {
			stack2 = append(stack2[:len(stack2)-1], stack2[len(stack2):]...)
		} else if string(i2) == "#"{
			continue
		}else {
			stack2 = append(stack2, string(i2))
		}
	}
	s = strings.Join(stack, "")
	t = strings.Join(stack2, "")
	if s == t {
		return true
	}
	return false

}
