func minimumSum(num int) int {
    a := []int{}
    for num >0 {
        a = append(a,num%10)
        num /= 10
    }
    sort.Ints(a)
    return a[0]*10 + a[1] *10 + a[2] + a[3]
}