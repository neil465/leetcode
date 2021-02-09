func restoreString(s string, indices []int) string {
	result := ""

	res := []string{}
	for i := 0; i < len(indices); i++ {
		res = append(res, "")
	}

	for i, index := range indices {

		res[index] = string(s[i])
		//fmt.Println(index, string(s[index]), res)
	}

	for _, re := range res {
		result += re
	}
	return result
}
