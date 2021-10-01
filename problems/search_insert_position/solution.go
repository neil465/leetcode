func searchInsert(nums []int, target int) int {
    if target == nums[0]{return 0}
    if target == nums[len(nums)-1]{return len(nums)-1}
    if target < nums[0]{return 0}
    if target > nums[len(nums)-1]{return len(nums)}
    l,r := 0,len(nums)-1
    mid := l + ((r - l) / 2)
    for l < r{
        if l+1 == r && nums[l]!= target && nums[r]!= target{return r}
        if nums[mid] > target{r = mid}
        if nums[mid] < target{l = mid}
        mid = l + ((r-l) / 2)
        if nums[mid] == target{
            return mid
        }
    }
    return mid 
}