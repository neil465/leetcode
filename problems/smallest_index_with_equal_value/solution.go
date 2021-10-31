func smallestEqual(nums []int) int {
    for i,j := range nums{
        if  i %10 == j{return i}
    } 
    return -1
}