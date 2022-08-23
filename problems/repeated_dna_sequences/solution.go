func findRepeatedDnaSequences(s string) []string {
    m := map[string]int{}
    for i := 0 ; i < len(s) - 9; i++ {
        m[s[i:i+10]] ++
    }
    arr := []string{}
    for i,j := range m {
        if j > 1 {
            arr = append(arr,i)
        }
    }
    return arr
}