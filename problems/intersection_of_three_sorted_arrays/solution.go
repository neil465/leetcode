func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
    arr := []int{}
    m,m2,m3 := make(map[int]bool),make(map[int]bool),make(map[int]bool)
    for _,i := range arr1{
        m[i] = true
    }
    for _,i := range arr2{
        m2[i] = true
    }
    for _,i := range arr3{
        m3[i] = true
    }
    for _,i := range arr1{
        if  m2[i] == true && m3[i] == true{
            arr = append(arr,i)
        }
    }
    return arr
    
}