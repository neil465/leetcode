func decrypt(code []int, k int) []int {
    g := make([]int, len(code))
    for i := 0 ; i < len(code); i++ {
        s := 0
        if k > 0 {
            start := i + 1
            for j := i + 1 ; j < i + k + 1; j ++ {
                if start == len(code) {
                    start = 0
                }
                
                s += code[start]
                start ++
            }
        }
        if k < 0 {
            start := i + k
            if start < 0 {
                start += len(code) 
            }
            for j := 0; j < -k ; j ++ {
                if start == len(code) {
                    start = 0
                }
                s += code[start]
                start ++
            }
        }
        g[i] = s
        
    }
    return g
}
