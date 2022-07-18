func maximumSum(nums []int) int {
    m := map[int][]int{}
    for _,i := range nums {
        a:=sum(i)
        m[a] = append(m[a],i)
    }
    max :=-1
    for _,i := range m {
        
        if len(i) >= 2 {
            sort.Ints(i)
            if a := i[len(i)-1] + i[len(i)-2] ; a > max {
            max = a
            }
        }
        
    }
    return max
}
func sum(i int) int {
    s := 0
    for i > 0 {
        s += i%10
        i /= 10
    }
    return s
}