

func halvesAreAlike(s string) bool {
    pre1 := []int{}
    curCount := 0

    for i := 0 ; i < len(s); i ++ {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
            curCount ++
        }
        pre1 = append(pre1, curCount)
    }
    pre2 := []int{}
    curCount = 0

    for i := len(s) - 1 ; i >= 0; i -- {
        if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
            curCount ++
        }
        pre2 = append([]int{curCount}, pre2...)
    }
    fmt.Println(pre1, pre2)
    return pre1[len(s)/2 - 1] == pre2[len(s)/2 ] 
}