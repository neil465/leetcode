
func maximum69Number(num int) int {
    sum := 0
    arr := []int{}
    
    // dividing up the number into a list of its digits
    for num > 0{
        arr = append(arr,num%10)
        num/=10
    }
    
    // reversing the list of digits , becasue when we made a list of all of the digits it was reversed
    l , r := 0 , len(arr)-1
    for l < r{ 
        arr[l] , arr[r] = arr[r] , arr[l]
        l++
        r--
    }
    
    // this is for finding the first digits that is a 6 and changing it to a nine
    for i2,i := range arr{
        if i == 6{
            arr[i2] = 9
            break
        }
    }
    // this is for finding how many zeros should the first number be multiplyed by
    nummultiplyer := 1
    for i := 0 ; i < len(arr)-1 ; i++{
        nummultiplyer *= 10
    }
    //and this is for remaking the list back into a array
    for _,i := range arr{
        sum += i*nummultiplyer
        nummultiplyer/=10
    }
    return sum
}
