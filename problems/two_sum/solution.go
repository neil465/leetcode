func twoSum(nums []int, target int) []int {
	mappy := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mappy[nums[i]]; ok {
			return []int{mappy[nums[i]],i}
		}else{
			mappy[(target-nums[i])] = i
		}
	}
	return []int{}
}