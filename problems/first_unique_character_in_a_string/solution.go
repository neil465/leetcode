func firstUniqChar(s string) int {
    arr := []string{}
    mapper := make(map[string]int)
    for _,i := range s{
        mapper[string(i)]++
    }
    for k,i := range mapper{
        if i == 1 {
            arr = append(arr,k)
        }
    }
    if len(arr) == 0{
        return -1
    }
    least := 1000000
    for _,i2 := range arr {
        for i,i3 := range s{
            if string(i3) == i2 && i<least{
                least = i
                break
            } 
        }
    }
    return least
}