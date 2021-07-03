func sortArrayByParity(A []int) []int {
    l,r := 0, len(A)-1
    for l <= r{
        if A[l]%2 ==0 {
            l++
        }else{
            if A[r]%2 != 0 {
                r--
            }else {
                A[l],A[r] = A[r],A[l]
            }
        }
    }
    return A
}
