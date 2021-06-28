
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ms := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(ms) > 0 && nums2[i] >= ms[len(ms)-1] {
			ms = ms[:len(ms)-1]
		}
		m[nums2[i]] = -1
		
		if len(ms) > 0 {
			m[nums2[i]] = ms[len(ms)-1]
		}
		ms = append(ms, nums2[i])
	}
	res := make([]int, len(nums1))

	for j, i := range nums1 {
		res[j] = m[i]
	}
	
	return res
}
