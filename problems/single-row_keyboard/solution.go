func calculateTime(keyboard string, word string) int {
    res := 0
    temp := ""
    alphabets := make(map[string]int)
    for j,i := range keyboard{
        alphabets[string(i)] = j
    }
    for _, i := range word{
        res += abs(alphabets[string(i)]-alphabets[temp])
        temp = string(i)
    }
    return res
}
func abs(i int) int{
    if i < 0{
        return i * -1
    }
    return i
}
