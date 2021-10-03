func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
    for j,i := range numbers{
        if _, ok := m[target-i]; ok{
            return []int{m[target-i]+1,j+1}
        }
        m[i] = j
    }
    return []int{}
}