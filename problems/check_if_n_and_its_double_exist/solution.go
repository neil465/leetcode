func checkIfExist(arr []int) bool {
    for g,i := range arr{
        if i%2 != 0{
            continue
        }
        for k,j := range arr{
            if j == i/2 && g != k{
                fmt.Println(j , i)
                return true
            }
        }
    }
    return false    
}