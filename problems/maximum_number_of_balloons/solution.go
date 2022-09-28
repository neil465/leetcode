func maxNumberOfBalloons(text string) int {
    m, min := [26]int{}, 100000
    for _, i := range text {
        m[i - 97]++
    }
    for _, i := range "balloon" {
        val := m[i - 97]
        if i == 'o' || i == 'l'{
            val /= 2
        }
        if val < min {
            min = val       
        }
    } 
    return min
}