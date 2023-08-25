func isConsecutive(nums []int) bool {
  sort.Ints(nums)
  for i := 0 ;  i < len(nums) -1 ; i ++ {
    if nums[i + 1] - nums[i] != 1 {
      return false
    }
  }
  return true
}