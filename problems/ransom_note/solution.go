func canConstruct(ransomNote string, magazine string) bool {
    m := make(map[string]int)
    max := len(magazine)
    if len(ransomNote) > len(magazine){
        max = len(ransomNote)
    }
    
    for i:= 0 ; i < max ; i++{
        if i<len(ransomNote){
            m[string(ransomNote[i])] --
        }
        if i<len(magazine){
            m[string(magazine[i])] ++
        }
    }
    for _,i := range m{
        if i <0 {
            return false
        }
    }
    return true
}