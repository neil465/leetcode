
func reverseVowels(s string) string {
	l,r := 0,len(s)-1
    res := []rune(s)
	for l<=r {
		for l <= r && isVowel(string(res[l])) == false {
            if l>r{
                return string(res)
            }
			l++
		}
		for isVowel(string(res[r])) == false {
            if l>r{
                return string(res)
            }
			r--
		}
		res[l] , res[r] = res[r],res[l]
        
		l++
		r--
	}
	return string(res)
}
func isVowel(target string) bool  {
	if target == "A" || target == "E" ||target == "I" || target == "O" ||target == "U" || target == "a" ||target == "e"||target == "i"||target == "o"||target == "u"{
		return true
	}
	return false
}