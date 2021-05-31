func countCharacters(words []string, chars string) int {
    res := 0
    fatty := make(map[string]int)
    for _,i2 := range chars{
        fatty[string(i2)]++
    }
   
    for _ , word := range words{
        flag := true
        fruity := make(map[string]int)
        for k,v := range fatty{
            fruity[k] = v
        }
        
        for _,letter := range word{
            if fruity[string(letter)] > 0 {
                fruity[string(letter)]--
            }else{
                
                flag = false
                break
            }
        }
        if flag{
            
            res+= len(word)
        }
        
    }
    return res
}