func minDeletionSize(strs []string) int {
    
    count := 0
    for i := 0 ; i < len(strs[0]) ; i++{
        bef := -1
        for _,j := range strs{
            if bef > int(j[i]){
                count++
                break
            }else{
                bef = int(j[i])
            }
        }
    }
    return count
}