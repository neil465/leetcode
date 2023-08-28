func maxSubsequence(nums []int, k int) []int {
  b := make([]int, len(nums))
  copy(b, nums)
  sort.Ints(nums)
  m := map[int]int{}
  for _, i := range nums[len(nums) - k:] {
    m[i] ++
  }
  a := []int{}
  for _, i := range b {
    if m[i] > 0 {
      a = append(a, i)
      m[i] --
    }
  }
  return a

}