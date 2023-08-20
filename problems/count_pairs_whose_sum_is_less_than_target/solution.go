func countPairs(nums []int, target int) int {
    p := 0
    for i := range nums {
      for j := i + 1; j < len(nums); j ++ {
        if nums[i]+ nums[j] < target {
          p++
        }
      }
    }
    return p
}