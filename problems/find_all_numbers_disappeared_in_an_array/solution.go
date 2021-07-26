
// func findDisappearedNumbers(nums []int) []int {
//     result := []int{}
//     m := make(map[int]int)
//     for _,i := range nums{
//         m[i]++
//     }
//     for i := 1 ; i <= len(nums);i++{
//         if m[i] == 0 {
//             result = append(result,i)
//         }
//     }
//     return result
    
// }
// func findDisappearedNumbers(nums []int) []int {
//     for i:= 0 ; i <len(nums) ; i++{
//         for nums[i] != i+1 && nums[i] != nums[nums[i]-1]{
//             nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
//         }
//     }
//     fmt.Println(nums)
//     result := []int{}
//     for i := 0 ; i < len(nums) ; i++{
//         if nums[i] != i+1{
//             result = append(result,i+1)
//         }
//     }
//     return result
// }
func findDisappearedNumbers(nums []int) []int {
    for i := 0 ; i <len(nums); i ++{
        if nums[abs(nums[i])-1] >0 {
            nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
        }
    }
    result := []int{}
    for i := 1 ; i <= len(nums) ; i ++{
        if nums[i-1] > 0 {
            result = append(result,i)
        }
    }
    return result
}
func abs(i int)int{
    if i < 0{
        return i * -1
    }
    return i 
}