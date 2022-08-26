func makeEqual(words []string) bool {
    m := make(map[string]int)
    for _,i := range words {
        for _, j := range i{
             m[string(j)]++   
        }
       
    }
    for _,i:= range m {
        if i % len(words) != 0 {
            return false
        }
    }
    return true
}