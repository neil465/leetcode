func balancedStringSplit(s string) int {
	res := 0
	L := 0
	R := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "R" {
			R ++
		}else {L++}
		if R==L {
			res++
		}
	}
	return res
}