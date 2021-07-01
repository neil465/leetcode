func maximumUnits(boxTypes [][]int, truckSize int) int {
    arr := []int{}
    for _,i := range boxTypes{
        for j := 0 ; j< i[0] ; j++{
            arr = append(arr,i[1])    
        }
    }
    res := 0
    sort.Ints(arr)
    count := 0
    for i := len(arr)-1 ; i >=0  ; i--{
        if count == truckSize {
            break
        }
        res+= arr[i]
        count++
    }
    return res
}