func sumOfDigits(nums []int) int {
    min := 100000
    for _,i := range nums{if i < min{min = i}}
    sum := 0
    for min>0{
        sum+=min%10
        min/=10
    }
    
    if sum & 1 == 1{
        return 0
    }
    return 1
}