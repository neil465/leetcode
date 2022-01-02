func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,j := range nums{

        if _, ok := m[target-j]; ok {
            return []int{m[target-j],i}
        }
        m[j] = i
    }
    return []int{}
    
}