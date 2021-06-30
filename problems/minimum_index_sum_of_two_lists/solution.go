func findRestaurant(list1 []string, list2 []string) []string {
    min := 100010121020120
    res := []string{}
    m := make(map[string]int)
    for i,j := range list1{
        m[j] = i+1
    }
    for i,j := range list2{
        if m[j] != 0{
            if m[j]-1 + i < min{
                min =  m[j]-1 + i
                res = []string{j}
            }
            if m[j]-1 + i == min{
                res = append(res,j)
            }
        } 
    }
    return findunique(res)
}
func findunique (arr []string) []string{
    flag := make(map[string]bool)
    var u []string
    for _, i := range arr {
        if flag[i] == false {
            flag[i] = true
            u = append(u, i)
        }
    }
    return u
}