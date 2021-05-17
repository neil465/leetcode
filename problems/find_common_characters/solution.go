func commonChars(A []string) []string {
    result := []string{}
    maps := []map[string]int{}
    for _,i2 := range A{
        m := make(map[string]int)
        for _,j := range i2{
            m[string(j)]++
        }
        maps = append(maps,m)
    }
   
    for k := range maps[0]{
        exist := true
        minimum := 1000000000000
        for i:=0;i<len(maps) ; i++{
            if val , ok := maps[i][k] ;!ok{
                exist = false
                break
            }else{
                if val < minimum {
                    minimum = val
                }
            }
        }
        if exist{
            for i := 0 ; i < minimum; i++{
                result = append(result,k)
            }
        }
       
    }
    
    return result
}