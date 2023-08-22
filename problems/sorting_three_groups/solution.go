type group struct {
  amounts [3]int
}

func minimumOperations(nums []int) int {

    groups := []group{group{}}
    
    for i := 0 ; i < len(nums); i ++ {
      last := [3]int{}
      last[0], last[1], last[2] = groups[i].amounts[0],groups[i].amounts[1], groups[i].amounts[2]
      last[nums[i] - 1] ++
      groups = append(groups, group{last})
    }
    // groups = groups[1:]
    min := 10000000
    for i := -1; i < len(nums); i ++ {
      for j := i  ; j < len(nums); j++ {
        sum := (groups[i + 1].amounts[1] + groups[i + 1].amounts[2]) + (groups[j + 1].amounts[0] - groups[i + 1].amounts[0] + groups[j + 1].amounts[2]- groups[i + 1].amounts[2]) + (groups[len(nums) ].amounts[0] - groups[j + 1].amounts[0] + groups[len(nums)].amounts[1]- groups[j + 1].amounts[1])
        if sum < min {
          min = sum
        }
      }
    }
    return min
}