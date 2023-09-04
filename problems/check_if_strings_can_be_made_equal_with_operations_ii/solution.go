func checkStrings(s1 string, s2 string) bool {
    odd := map[byte]int{}
    even := map[byte]int{}
    for i := range s1 {
        if i % 2 == 0 {
            even[s1[i]] ++
        } else {
            odd[s1[i]] ++
        }
    }
    for i := range s2 {
        if i % 2 == 0 {
            if even[s2[i]] == 0 {
                return false
            }
            even[s2[i]] --
        } else {
            if odd[s2[i]] == 0{
                return false
                }
            odd[s2[i]] --
        }
        
    }
    return true
}