func asteroidsDestroyed(mass int, asteroids []int) bool {
    arr := []int{}
    for _,i := range asteroids{
        if i > mass{
            arr = append(arr,i)
        }else{
            mass += i            
        }
 
    }
    sort.Ints(arr)
    for _,i := range arr{
        if i > mass{
            return false
        }
        mass += i
    }
    return true
}