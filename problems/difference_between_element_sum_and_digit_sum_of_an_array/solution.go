func differenceOfSum(nums []int) int {
    esum, dsum := 0, 0 
    for _, i := range nums {
        esum += i
        
        for i > 0 {
            dsum += i %10
            i /= 10    
        }
    }
    return esum - dsum
}