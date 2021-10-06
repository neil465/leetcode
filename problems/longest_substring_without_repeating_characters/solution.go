func lengthOfLongestSubstring(s string) int {
    max := 0 
    var ind [128]int
    
    for i,j := 0,0 ; j < len(s) ; j++{
        if ind[s[j]] > i{
            i = ind[s[j]]
        }
        
        if j-i+1 > max{
            max = j-i+1
        }
        ind[s[j]] = j+1
    }
    return max
}