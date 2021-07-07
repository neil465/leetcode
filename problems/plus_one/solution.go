func plusOne(digits []int) []int {
    t := []int{1}

   for i := len(digits)-1;i>=0 ; i--{
        if digits[i] < 9{
            digits[i] = digits[i]+1
            break
        }else if i == 0{
            digits[0] = 0
            digits = append(t, digits...)
            break
        }else{
            digits[i] = 0
        }
    }

    return digits
}
