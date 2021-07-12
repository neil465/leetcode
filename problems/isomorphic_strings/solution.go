func isIsomorphic(s string, t string) bool {
    m1 ,m2 := make(map[rune]int),make(map[rune]int)
    count := 1
    for _,i := range s{
        if m1[i] == 0{
            m1[i] = count
            count++
        }
    }
    count = 1
    for _,i := range t{
        if m2[i] == 0{
            m2[i] = count
            count++
        }
    }
   
    
    arr2 := make([]int,len(s))
    arr1 := make([]int,len(s))
    
    for j,i := range s{
        arr1[j] = m1[i]
    }
    for j,i := range t{
        arr2[j] = m2[i]
    }
    return reflect.DeepEqual(arr1, arr2)
}