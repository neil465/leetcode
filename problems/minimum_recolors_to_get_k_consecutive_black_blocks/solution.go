func minimumRecolors(blocks string, k int) int {
    min := 10000
    for i := 0 ; i < len(blocks) -(k-1 ) ; i++{
        sum := 0 
        for j := i ; j < i+k;j++{
            if string(blocks[j]) == "W" {
                sum++
            }
        }
        if sum < min {
            min= sum
        }
    }
    return min
}