
func maximum69Number(num int) int {
    sum := 0
    arr := []int{}
    for num > 0{
        arr = append(arr,num%10)
        num/=10
    }
    l , r := 0 , len(arr)-1
    for l < r{
        arr[l] , arr[r] = arr[r] , arr[l]
        l++
        r--
    }
    for i2,i := range arr{
        if i == 6{
            arr[i2] = 9
            break
        }
    }
    nummultiplyer := 1
    for i := 0 ; i < len(arr)-1 ; i++{
        nummultiplyer *= 10
    }
    for _,i := range arr{
        sum += i*nummultiplyer
        nummultiplyer/=10
    }
    return sum
}
