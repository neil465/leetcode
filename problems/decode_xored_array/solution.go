func decode(encoded []int, first int) []int {
    arr := []int{first}
    for i := 0 ; i < len(encoded); i++{
        fmt.Println(i)
        for j:= 0 ; j <= 1000000000000 ; j++{
            if arr[len(arr)-1]^j == encoded[i]{
                arr = append(arr,j)
                break
            }
        }
    }
    return arr
}