func search(nums []int, target int) int {
    l,r := 0,len(nums)-1
    mid := l + ((r - l) / 2)
    if nums[0] == target{return 0}
    if nums[len(nums)-1] == target{return len(nums)-1}
    if nums[mid]  == target{return mid}
    for l<r-1{
        
        if nums[mid] < target{l=mid}
        if nums[mid]  > target{r=mid}
        mid = l + ((r - l) / 2)
        if nums[mid]  == target{return mid}
    }
    return -1

}