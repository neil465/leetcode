func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i,j := range nums{

        if _, found := m[target-j]; found {
            return []int{m[target-j],i}
        }
        m[j] = i
    }
    return []int{}
    
}