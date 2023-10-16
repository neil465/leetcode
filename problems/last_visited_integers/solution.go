func lastVisitedIntegers(words []string) []int {
    a := []int{}
    res := []int{}
    curPointer := 0 
    for i := range words {
        if words[i] == "prev" {
            if i == 0 || words[i] != words[i - 1] {
                curPointer = 0
            }
            if len(a)  - curPointer <= 0 {
                res = append(res, -1)
            } else {
                res = append(res, a[len(a) - 1 - curPointer])
            }
            curPointer ++
        } else {
            k, _:= strconv.Atoi(words[i])
            a = append(a, k)
        }
    }
    return res
}