func countQuadruplets(nums []int) int {
    count := 0
    for i:= 0 ; i < len(nums)-3 ; i++{
        for j := i+1 ; j<len(nums) ; j++{
            for k := j+1 ; k<len(nums) ; k++{
                for l := k+1 ; l<len(nums) ; l++{
                    if nums[i]+nums[j]+nums[k] == nums[l]{ 
                        count ++
                    }
                }
            }
        }
    }
    return count
}