func myAtoi(s string) int {
    
    flag := false
    sum := 0
    arr := []int{}
    multiplier := 1

    s = strings.Trim(s," ") 

    for j,i := range s{
        if j != len(s)-1 && (s[j:j+2] == "+-" || s[j:j+2] == "-+"){
            return 0
        }
        switch i {
            
            case ' ':
                if flag{
                    break
                } 
                continue
            case '-':
                if j != len(s)-1 && s[j+1] == ' '{
                    return 0
                }   
                if j !=0 {
                    break
                }
                multiplier = -1
                continue
            case '+':
            
                if j != len(s)-1 && s[j+1] == ' '{
                    return 0
                }   
                if j !=0 {
                    break
                }
                continue
                 
        }
        if unicode.IsDigit(i){
            
            flag = true
            arr = append(arr,int(i)-48)
        }else{
            break
        }    
    }
    for len(arr) > 0 {
        if len(arr)-1 > 18 && arr[0] != 0{
            if multiplier == 1{
                return 2147483647
            }
            if multiplier == -1{
                return -2147483648
            }
        } 
        if arr[0] * int(math.Pow(10, float64(len(arr))-1))* multiplier > 2147483647{
            return 2147483647
        }
        if arr[0] * int(math.Pow(10, float64(len(arr))-1))* multiplier < -2147483648{
            return -2147483648
        }
        fmt.Println(int(math.Pow(10, float64(len(arr))-1))) 
        sum += arr[0] * int(math.Pow(10, float64(len(arr))-1))
        arr = arr[1:]
    }
    if sum* multiplier > 2147483647{
        return 2147483647
    }
    if sum* multiplier < -2147483648{
        return -2147483648
    }
    return sum*multiplier
}


