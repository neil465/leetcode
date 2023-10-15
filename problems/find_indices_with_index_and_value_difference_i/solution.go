func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    for i := range nums {
        for j := range nums {
            if abs(i - j) >= indexDifference && abs(nums[i] - nums[j]) >= valueDifference {
                return []int{i, j}
            } 
        }
    }
    return []int{-1,-1}
}
func abs(i int) int {
    if i > 0 {
        return i
    }
    return -i
}