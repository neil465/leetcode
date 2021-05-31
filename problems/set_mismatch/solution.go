func findErrorNums(nums []int) []int {
    mapy := make(map[int]int)
    for _,i := range nums{
        mapy[i]++
    }
    doublenum := 0
    num := 0
    
    for i := 0 ; i < len(nums)+1 ; i++{
        if mapy[i] == 0{
            num = i
        }
        if mapy[i] == 2{
            doublenum = i
        }
    }
    return []int{doublenum,num}
}