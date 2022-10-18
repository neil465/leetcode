func sumOfNumberAndReverse(num int) bool {

    for i := num/2 ; i <= num; i++ {
        rev, temp :=  0, i
        for temp > 0 {
            rev, temp = rev * 10 + temp%10 , temp/10
        }
        if i + rev == num {
            return true
        } 
    }
    return false
}