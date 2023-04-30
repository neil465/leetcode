func isWinner(player1 []int, player2 []int) int {
    c := 0
    r1, r2 := 0, 0
    for _, i := range player1 {
        
        if c > 0 {
            c -- 
            r1 += i
        }
        if i == 10 {
            c = 2
        } 
        r1 += i
        
    }
    c= 0
    fmt.Println(r1)
    for _, i := range player2 {
        if  c > 0 {
            c -- 
            r2 += i
        }
        if i == 10 {
            c = 2
        } 
        r2 += i
        
        
    }
    fmt.Println(r2)
    if r1 > r2 {
        return 1
    } else if r1 < r2 {
        return 2
    }
    return 0
    
}