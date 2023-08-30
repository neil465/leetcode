func searchRange(nums []int, target int) []int {
  if len(nums) == 0 {
    return []int{-1, -1}
  }
    mid, s, e := len(nums) / 2, 0 ,len(nums) - 1
    id1, id2 := -1,-1 
    for  s <= e {
      mid = (s + e) / 2
      if nums[mid] < target {
        s = mid + 1
      } else {
        e = mid - 1
      }
      if nums[mid] == target {
        id1 = mid
      }
      
    }
    mid, s, e = len(nums) / 2, 0 ,len(nums) - 1
    for  s <= e {
      mid = (s + e) / 2
      if nums[mid] <= target {
        s = mid + 1
      } else {
        e = mid - 1
      }
      if nums[mid] == target {
        id2 = mid
      }
      
    }
    
    return []int{id1, id2 }
}