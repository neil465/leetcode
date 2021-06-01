func longestCommonPrefix(strs []string) string {
    res := ""
    sort.Strings(strs)
    for i := 0 ; i < len(strs[0]) ; i++{
        val:= strs[0]
        if i+1 < len(strs[0]){
            val = strs[0][:i+1]
        }
        for _,j := range strs{
            val2 := j[:i+1]
            if val != val2{
                return res
            }
        }
        res = val
    }
    return res
}