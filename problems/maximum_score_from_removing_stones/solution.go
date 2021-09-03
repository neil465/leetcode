// func maximumScore(a int, b int, c int) int {
//     solitaire := []int{a,b,c}
//     sort.Ints(solitaire)
//     score := 0
//     for solitaire[1] != 0 && solitaire[2] != 0{
//         score++
//         solitaire[1]--
//         solitaire[2]--
//         sort.Ints(solitaire)
//     }
//     return score
// }

func maximumScore(a int, b int, c int) int {
    if a > b{return maximumScore(b,a,c)}
    if b > c{return maximumScore(a,c,b)}
    if a+b <=c{return a+b}
    return c+(a+b-c)/2
}