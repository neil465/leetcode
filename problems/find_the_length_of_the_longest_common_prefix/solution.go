func longestCommonPrefix(arr1 []int, arr2 []int) int {
    m := map[int]int{}
    
    for _, i := range arr1 {
        v := i
        
        for v > 0 {
            if v == 0 {
                continue
            }
            m[v] = i
            v /= 10
        }

    }
    maxL := 0
     for _, i := range arr2 {
        v := i
        
        for v > 0 {
            if _, ok := m[v]; ok {
                // fmt.Println(m, v)
                maxL = max(maxL, v)
                break
            }
            
            v /= 10
        }

    }
    
    v := 0 
    
    for maxL > 0 {
        maxL /= 10
        v ++
    }
    
    return v
}