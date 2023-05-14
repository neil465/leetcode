func circularGameLosers(n int, k int) []int {
    index :=  1
    m := map[int]bool{1 : true}
    c := 1
    for {
        index += k *c
        mod := index % n
        if mod == 0 {
            mod = n
        }
        if m[mod] {
            a := []int{}
            for i := 1 ; i <= n ; i++ {
                if !m[i]  {
                    a = append(a, i)
                }
            }
            
            return a
            
        }
        c ++
        m[mod] = true
    }
    return []int{}
}