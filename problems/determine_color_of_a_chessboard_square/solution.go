func squareIsWhite(coordinates string) bool {
    q1 := coordinates[0] - 96
    q2 ,_:= strconv.Atoi(string(coordinates[1]))
    if q1 % 2 == 0{
        if q2 %2 == 0{
            return false
        }
        return true
    }else {
        if q2 %2 == 0{
            return true
        }
        return false
    }
    return true
}