func longestAlternatingSubarray(nums []int, threshold int) int {
    if len(nums) <= 1{
        if len(nums) ==1 {
            if nums[0] % 2 != 0 || nums[0] > threshold{
                return 0
            }
        }
        return len(nums)
    }
    maxlen := 0
    k := true
    for l := 0 ; l < len(nums); l ++ {
        for r := l  ; r < len(nums); r++ {
            if nums[l] % 2 != 0 || nums[l] > threshold{
                continue
            }
            b := false
            for i := l ; i < r; i ++ {
                if nums[i] % 2 == nums[i + 1] % 2 || nums[i] > threshold || nums[i + 1] > threshold{
                    b = true
                    break
                }
            }
            if b {
                continue 
            }

            if r - l + 1 > maxlen {
                maxlen = r - l + 1

                k = false

            }
            
        }
    }
    if k {

        return 0
    }
    return maxlen
}