func minBitFlips(start int, goal int) int {
    pointer := 1 
    sum := 0
    for i := 0 ; i < 30 ; i++ {
        if start & pointer != goal & pointer {

            sum++
        }
        pointer *= 2
    }
    return sum
}