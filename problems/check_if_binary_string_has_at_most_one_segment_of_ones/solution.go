func checkOnesSegment(s string) bool {
    flag := false
    for _,i2 := range s{
        if i2 == '0' {
            flag = true
        }
        if flag == true && i2 == '1'{
            return false
        }
    }
    return true
}