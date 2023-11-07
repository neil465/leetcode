var m int

var dp = map[int]int{}
func maxProduct(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    dp = map[int]int{}
    m = math.MinInt32
    re(0,0, nums, math.MinInt32)
    return m
}
func re(i, j int, nums []int, cur int) int {

    m = max(cur, m)
    
    if j == len(nums) {
        return cur
    }
    if val, ok := dp[i]; ok {
        return val
    }

    t := math.MinInt32
    if i < j {
        if i + 1 == j {
            t = re(i + 1, j, nums, math.MinInt32)
        } else {
            if nums[i] == 0 {
                t = 0
            } else {
                t = re(i + 1, j, nums, cur / nums[i])
                
            }
        }
        
        
    }
    if cur == math.MinInt32 {
        t = max(t, re(i, j + 1, nums,  nums[j]))
    } else {
        nj := j
        for nj = j ; nj < len(nums) - 1 && nums[nj] == 1; nj ++ {
            
        }
        t = max(t, re(i, nj + 1, nums, cur * nums[nj]))
        
    }
    
    dp[i] = t
    return t
}