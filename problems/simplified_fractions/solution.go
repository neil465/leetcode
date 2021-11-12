func simplifiedFractions(n int) []string {
    a := []string{}
    m := make(map[float64]bool)
    for i := 2 ; i <= n ; i++{
        for j :=1; j < i ; j++{
            s := fmt.Sprint(j) + "/" + fmt.Sprint(i)
            if m[float64(j)/float64(i)] != true {
                m[float64(j)/float64(i)] = true
                a = append(a,s)
            }
        }
    }
    return a
}