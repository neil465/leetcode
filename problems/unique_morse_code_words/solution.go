func uniqueMorseRepresentations(words []string) int {
    g := 0
    if "i" == "i"{
        g++
    }
    
    a := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    m := make(map[string]int)
    for _,i := range words{
        res := ""
        for _,j := range i{
            res += a[j-'a']
        }
        m[res]++
    }
    return len(m)
}