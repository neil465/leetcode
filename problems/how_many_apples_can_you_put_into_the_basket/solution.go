func maxNumberOfApples(arr []int) int {
    sort.Ints(arr)
    sum := 0
    
    fmt.Println(arr)
    amount := len(arr)
    for _,i := range arr{
        sum += i
    }
    for len(arr)>0 && sum > 5000{
        sum -= arr[len(arr)-1]
        arr = arr[:len(arr)-1]
        amount--
    }
    fmt.Println(arr)
    return amount
}