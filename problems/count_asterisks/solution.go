func countAsterisks(s string) int {
    result, count := 0, 0 
    for _, i := range s {
        if i == '|' {
            count ++
        }
        if count % 2 == 0 && i == '*' {
            result ++
        }
    }
    return result
}