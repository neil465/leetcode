func stringMatching(words []string) []string {
    res := []string{}
    sort.Slice(words,func (i,j int)bool{
        return len(words[i])<len(words[j])
    })
    for j,i := range words{
        for k := j ; k< len(words) ; k++{
            if k != j{
               
                if strings.Contains(words[k],i){
                
                    res = append(res,i)
                    break
                }
            }
        }
    }
  
    return res
}