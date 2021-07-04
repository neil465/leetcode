func findClosestElements(arr []int, k int, x int) []int {
    sort.Slice(arr,func(i,j int)bool{
        if abs(arr[i] , x) < abs(arr[j] , x) ||( abs(arr[i] , x) == abs(arr[j] , x) && arr[i]<arr[j]){
            return true
        }
        return false
    })
    res := arr[:k]
    sort.Ints(res)
    return res
}

func abs(i,j int)int{
    if i-j < 0{
        return (i-j)*-1
    }
    return i-j
}