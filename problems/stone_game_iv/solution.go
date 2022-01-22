var memo = map[int]bool{} 
func winnerSquareGame(n int) bool {
    if memo[n]{
        return memo[n]
    }
    if n <= 0 {
        return false
    }
    for  j := 1; j * j < n + 1; j++{


        if !winnerSquareGame(n-(j*j)){
            memo[n] = true
            return true
        }


    }
    return false
}
