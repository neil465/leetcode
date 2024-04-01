func minimumSubarrayLength(nums []int, k int) int {
    l := math.MaxInt32
    arr := [32]int{}

    i := 0 

    cur := 0

    for j := range nums {
        cur |= nums[j] 

        for iter := 0 ; iter < 32; iter ++ {
            if nums[j] & (1 << iter) > 0 {
                arr[iter] ++
            }
        }

        for i = i ; i <= j && cur >= k; i ++ {
            l = min(l, j - i + 1)
            ncur := 0
            for iter := 0 ; iter < 32; iter ++ {
                if nums[i] & (1 << iter) > 0  {
                    arr[iter] --
                }
                if arr[iter] > 0  {
                    ncur |= 1 << iter
                }
            }
            cur = ncur
        }
        
    }

    
    if l == math.MaxInt32 {
        return -1
    }
    return max(l, 1)
}
