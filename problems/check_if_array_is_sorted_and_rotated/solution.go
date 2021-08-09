func check(nums []int) bool {
    
     c:= make([]int, len(nums))
    copy(c,nums)
    sort.Ints(c)
    for i := 0 ; i < len(c) ; i++{
        last := c[len(c)-1]
        c = c[:len(c)-1]
        c = append([]int{last},c...)
        
        if reflect.DeepEqual(nums,c){
            fmt.Println(nums,c)
            return true
        }
    }
    return false
}