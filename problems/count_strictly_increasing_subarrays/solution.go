func countSubarrays(nums []int) int64 {
    i,j := 0,0
    sum := int64(0) 
    for j < len(nums)-1{
        if nums[j] >= nums[j+1] {
            maxSize := j-i+1
            sum += int64((maxSize) * (maxSize + 1) / 2)
            i = j+1
        }
        j++
    }
    maxSize := j - i + 1
    sum += int64((maxSize) * (maxSize + 1) / 2)
    return sum
}