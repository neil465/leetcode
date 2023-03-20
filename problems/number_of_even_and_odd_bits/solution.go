func evenOddBit(n int) []int {
    pointer := 1
    even, odd := 0, 0
    for i := 0 ; i < 40; i++ {
        if n & pointer == 0 {
            if i % 2 == 0{
                even++
            }else {
                odd++
            }
        }
        pointer *= 2
    }
    return []int{20 - even, 20 -odd}
}