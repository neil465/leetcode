func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }
    start, finish, arr := nums[0],nums[0], []string{}

    for i := 0 ; i < len(nums) - 1 ; i++{
        if nums[i+1] - nums[i]  < 2 {
            finish ++
        } else {
            if start == finish {
                arr = append(arr, fmt.Sprint(start))
            } else {
                arr = append(arr, fmt.Sprint(start) + "->" + fmt.Sprint(finish))          
            }
            start = nums[i + 1]
            finish = nums[i + 1]
        }
    }
    if start == finish {
        arr = append(arr, fmt.Sprint(start))
    } else {
        arr = append(arr, fmt.Sprint(start) + "->" + fmt.Sprint(finish))          
    }
    return arr
}