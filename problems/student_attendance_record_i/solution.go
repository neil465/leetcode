func checkRecord(s string) bool {
    absentTotal := 0
    for j,i := range s{
        if string(i) == "A"{
            if absentTotal == 1 {
                return false
            }else{
                absentTotal++
            }
        }
        if string(i) == "L" && j+2 <= len(s)-1{
            if string(s[j+1]) == "L" && string(s[j+2]) == "L"{
                return false
            }
        }
        
    }
    return true
}