func secondHighest(s string) int {
    m := make(map[int]int)
    a := []int{}
    for _,i := range s{
        if !unicode.IsLetter(i){
            g,_ := strconv.Atoi(string(i))
            if m[g] == 0{
                a = append(a, g)
                m[g]++
            }
        }
    }
    if len(a) <2{
        return -1
    }
    sort.Ints(a)
    return a[len(a)-2]
}