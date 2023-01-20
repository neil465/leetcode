func countDigits(num int) int {
    divisibles := 0 

    temp := num

    for temp > 0 {
        a := temp%10
        if num % a == 0 {
            divisibles ++
        } 
        temp /= 10
    }
    return divisibles 
}