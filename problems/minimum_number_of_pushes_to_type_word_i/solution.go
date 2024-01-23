func minimumPushes(word string) int {

    if len(word) <= 8 {
        return len(word)
    }
    if len(word) <= 16 {
        return 8 + (len(word) - 8) * 2
    } 
    if len(word) <= 24 {
        return 24 + (len(word) - 16) * 3
    }
    return 48 + (len(word) - 24) * 4
}