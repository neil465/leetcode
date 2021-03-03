func isValid(s string) bool {
	stack := []string{}
	for _, i2 := range s {
		arr := []string{"[","(","{"}
		arr2 := []string{"]",")","}"}
		//if string(i2) == "(" {
		//	stack = append(stack, string(i2))
		//}
		//if string(i2) == ")" {
		//	if len(stack) > 0 && stack[len(stack)-1] == "(" {
		//		stack = stack[:len(stack)-1]
		//	} else {
		//		stack = append(stack, string(i2))
		//	}
		//
		//}
		//if string(i2) == "[" {
		//	stack = append(stack, string(i2))
		//}
		//if string(i2) == "]" {
		//	if len(stack) > 0 && stack[len(stack)-1] == "[" {
		//		stack = stack[:len(stack)-1]
		//	} else {
		//		stack = append(stack, string(i2))
		//	}
		//
		//}
		for i, s2 := range arr {
			if string(i2) == s2 {
					stack = append(stack, string(i2))
				}
				if string(i2) == arr2[i] {
					if len(stack) > 0 && stack[len(stack)-1] == s2 {
						stack = stack[:len(stack)-1]
					} else {
						stack = append(stack, string(i2))
					}
				
				}
		}

	}
	fmt.Println(stack)
	return len(stack) < 1
}
