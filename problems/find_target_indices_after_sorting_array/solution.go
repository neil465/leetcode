func targetIndices(nums []int, target int) []int {
    smallTarget, equalTarget := 0,0
    for _,i := range nums {
        if i < target {
            
            smallTarget++
        }else if i == target { 
            equalTarget++
        }
    }
    fmt.Println(smallTarget, equalTarget)
    arr := []int{}
    for i := 0 ; i< equalTarget ;i++{
        arr = append(arr,smallTarget)
        smallTarget++
    }
    return arr
}