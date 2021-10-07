func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2){return false}
    i,j := 0,len(s1)-1
    m:= make(map[string]int)
    for _,i := range s1{
        m[string(i)]++
    }
    for j < len(s2){
         g:= make(map[string]int)
      
    
        for index, element  := range m{        
            g[index] = element
        }
        
        for _,i := range s2[i:j+1]{
            if g[string(i)] == 1{
                delete(g,string(i))
            }else{
                g[string(i)]--    
            }
            
        }
      
        if len(g) == 0{return true}
        i++
        j++
    }
    return false
}