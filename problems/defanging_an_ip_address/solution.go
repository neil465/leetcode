
func defangIPaddr(address string) string {
	result := ""
	for _, i2 := range address {
		if string(i2) == "." {
			result += "[.]"
		} else {
			result += string(i2)
		}
	}
	return result
}
