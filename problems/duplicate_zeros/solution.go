func duplicateZeros(arr []int)  {
    for i:= 0; i < len(arr); i++ {
        if arr[i] == 0 {
            copy(arr[i+1:], arr[i :])
            
            i++
        }
    }
}