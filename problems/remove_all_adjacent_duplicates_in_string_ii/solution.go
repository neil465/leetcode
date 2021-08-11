func removeDuplicates(s string, k int) string {
     for i := 0 ; i < len(s)-k+1 ; i++{
         flag := true
         for d := i+1 ; d < i+k ; d++{
             if s[i] != s[d]{
                flag = false
             }
         }
        if flag {
            s = s[:i] + s[i+k:]
            i =-1
        }
    }
    return s
}