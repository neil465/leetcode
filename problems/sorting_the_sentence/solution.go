func sortSentence(s string) string {
     
    splitstring:= strings.Split(s , " ")
    arr := make([]string , len(splitstring))
    for i , i2 := range splitstring{
        g,_ :=strconv.Atoi(string(i2[len(i2)-1]))
        fmt.Println(g)
        arr[g-1] = splitstring[i][:len(i2)-1]
        fmt.Println(arr,i2)
    }
    s = strings.Join(arr , " ")
    return s
}