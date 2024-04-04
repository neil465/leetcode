func maxDepth(s string) int {
	max := 0 
	stack := []string{}
	for _, i2 := range s {
		if string(i2) == "(" || string(i2) == ")" {
			if string(i2) == "(" {
				stack = append(stack,"(")
				if len(stack)>max {
					max = len(stack)
				}
			}else{
				stack = stack[:len(stack)-1]
			}
		}
	}
	return max
}
