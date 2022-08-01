func minimumOperations(nums []int) int {
    m:=map[int]int{}
    for _,i:=range nums{
        if i>0{
            m[i]++
        }
    }
    return len(m)
}