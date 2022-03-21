func partitionLabels(s string) []int {
    m, result := map[rune]int{}, []int{}
    
    for i, j := range s {
        m[j] = int(math.Max(float64(i),float64(m[j])))
    }
    last , start := 0,0
    for i := range s {
        
        last = int(math.Max(float64(last),float64(m[rune(s[i])])))
        
        if i == last {
            result = append(result, last - start + 1)
            start = last + 1
        }
        
    }
    return result
}