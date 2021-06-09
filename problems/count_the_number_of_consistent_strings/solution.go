func countConsistentStrings(allowed string, words []string) int {
    res := 0
    for _,j := range words{
        flag := false
        for _, k := range j{
            flag = false
            for _, i := range allowed{
                if i == k {
                    flag = true
                   
                }
            }
            if flag == false {
                break
            }
            
            
        }
        if flag {
            res++
        }
    
    }
    return res
    
}
