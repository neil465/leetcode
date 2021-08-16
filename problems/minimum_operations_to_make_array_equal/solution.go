func minOperations(n int) int {
    sum,count := 0,0
    for i := 1 ; i < n ; i++{
        if i % 2 != 0{count++}
        sum+= count
    }
    return sum
}