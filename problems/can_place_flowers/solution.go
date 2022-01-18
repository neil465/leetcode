func canPlaceFlowers(flowerbed []int, n int) bool {
    flowerbed = append(flowerbed,0)
    flowerbed = append([]int{0},flowerbed...)
    for i := 1 ; i < len(flowerbed)-1 ; i++{
        if flowerbed[i+1] + flowerbed[i] + flowerbed[i-1] == 0{
            i++
            n--    
        }
        
    }
    return n <= 0
}