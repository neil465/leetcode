func similarPairs(words []string) int {
    m := map[[26]bool]int{}

    sum := 0 

    for _, i := range words {
        freq := [26]bool{}
        for _, a := range i {
            freq[int(a) - 97] = true
        }
        
        sum += m[freq]
        m[freq] ++
    }
    return sum
}