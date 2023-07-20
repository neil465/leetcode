func numSteps(s string) int {
    g := []rune(s)
    q := 0
    for len(g) > 1 {
        if g[len(g) - 1] == '1' {
            
            ind := len(g) - 1
            for ind >= 0 && g[ind] == '1'{

                g[ind] = '0'
                
                ind --
            }
            // fmt.Println(ind, string(g), "sig1")
            if ind == -1  {
                g = []rune("1" + string(g))
            } else {
                g[ind] = '1'
            }

        
        } else {

            g = g[:len(g) - 1]
        }

        q ++
    }
    return q


}