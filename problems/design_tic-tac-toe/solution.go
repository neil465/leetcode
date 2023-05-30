type TicTacToe struct {
    a [][]int
}


func Constructor(n int) TicTacToe {
    a := make([][]int, n)
    for i := range a {
        a[i] = make([]int, n)
    }
    return TicTacToe{a}
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    this.a[row][col] = player
    for i := 0 ; i < len(this.a); i ++ {
        v := this.a[i][0]
        a := true
        k := this.a[0][i]
        g := true
        for j := 0; j < len(this.a); j ++ {
            if k != this.a[j][i] {
                g = false
            }
            if v != this.a[i][j] {
                a = false
            }
        }
        if a {
            return v
        }
        if g {
            return k
        }
    }
    v := this.a[0][0]
    b := true
    for i := 0 ; i < len(this.a); i ++ {
        if this.a[i][i] != v {
            b = false
            break
        }
    }
    if b {
        return v
    }
    v = this.a[0][len(this.a) -1]
    b = true
    for i := 0 ; i < len(this.a); i ++ {
        if this.a[i][len(this.a) - 1 - i] != v {
            b = false
            break
        }
    }
    if b {
        return v
    }
    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */