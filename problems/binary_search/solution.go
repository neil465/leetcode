
func search(nums []int, target int) int {
	low, mid, hi := 0, 0, len(nums)-1
    
	for low <= hi {
		mid = low + (hi-low)/2
		if  nums[mid] == target {
			return mid
		}else if  nums[mid] < target {
			low = mid+1
		} else {
			hi = mid - 1
		}
	}
    
	return -1
}
