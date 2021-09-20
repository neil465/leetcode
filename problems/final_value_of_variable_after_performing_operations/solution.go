func finalValueAfterOperations(operations []string) int {
    count := 0
    for _,i := range operations{
        if i == "--X" || i == "X--"{count--}
        if i == "++X" || i == "X++"{count++}
    }
    return count
}