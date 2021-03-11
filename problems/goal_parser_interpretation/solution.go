func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); i++ {
		if string(command[i]) == "("  {
			if string(command[i+1]) != "a" {
				i++
				res+="o"
			}else {
				i+=3
				res+="al"
			}
		}else {
			res+="G"
		}
	}
	return res
}