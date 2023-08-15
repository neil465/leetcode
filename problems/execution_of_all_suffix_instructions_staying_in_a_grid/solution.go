func executeInstructions(n int, startPos []int, s string) []int {
    res := []int{}

    for j := range s {
        got := j - 1
        cp := []int{startPos[0], startPos[1]}
        for i := j ; i < len(s); i ++ {
            
            if s[i] == 'R' {
                cp[1] ++
            } else if s[i] == 'L' {
                cp[1] --
            } else if s[i] == 'D' {
                cp[0] ++
            } else if s[i] == 'U' {
                cp[0] --
            }
            if cp[0] < 0 || cp[0] >= n || cp[1] < 0 || cp[1] >= n {

                break
            }
            got = i
            
        }
        res = append(res, got - j + 1 )


    }
    return res
    
}