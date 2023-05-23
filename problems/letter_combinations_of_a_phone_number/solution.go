var m  = map[string]string{
    "2" : "abc",
    "3" : "def",
    "4" : "ghi",
    "5" : "jkl",
    "6" : "mno",
    "7" : "pqrs",
    "8" : "tuv",
    "9" : "wxyz"}
func letterCombinations(digits string) []string {
    a := helper(digits,"")    
    if a[0] == ""{
        return []string{}
    }
    return a
}
func helper(digits string , str string) []string{
    if len(digits) == 0 {
        return []string{str}
    }
    
    res := []string{}
    
    for _,i := range m[string(digits[0])]{
        res = append(res,helper(digits[1:],str+string(i))...)    
    }
    
    return res
}