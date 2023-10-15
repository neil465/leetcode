func findIndices(nums []int, indexDifference int, valueDifference int) []int {

    m := map[int]int{}
    for i := range nums {
        m[nums[i]] = i
    }
    maxCur := math.MinInt32
    preSum := make([]int, len(nums))
    for i := len(nums) -1 ; i>= 0; i-- {
        preSum[i] = max(nums[i ], maxCur)
        maxCur = max(nums[i], maxCur)
    }
    maxCur = math.MaxInt32
    preSum2 := make([]int, len(nums))
    for i := len(nums) -1 ; i>= 0; i-- {
        preSum2[i] = min(nums[i ], maxCur)
        maxCur = min(nums[i], maxCur)
    }

    for i := 0; i < len(nums) ;i++ {
        if indexDifference + i < len(preSum) {
            if abs(preSum[indexDifference + i] - nums[i]) >= valueDifference  {
                
                return []int{m[preSum[indexDifference + i]], i} 
            }
            if abs(preSum2[indexDifference + i] - nums[i]) >= valueDifference {
                return []int{m[preSum2[indexDifference + i]], i} 
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
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}