func isStrobogrammatic(num string) bool {
    flips := map[string] string {
        "6" : "9",
        "8" : "8",
        "9" : "6",
        "0" : "0",
        "1" : "1",
    }      
    for i := 0 ; i < len(num); i++{
        if flips[string(num[i])] != string(num[len(num) - i-1]) {
            return false
        }
    }
    return true
}