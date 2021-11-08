// func subsets(nums []int) [][]int {
//     res := [][]int{}
//     for i := 0 ; i < int(math.Pow(2,float64(len(nums)))); i++{
//         num := i
     
//         temp := []int{}
//         pointer := 0
//         for num != 0{
//             if num & 1 == 1{
//                 temp = append(temp,nums[pointer])
//             }
//             pointer++
//             num = num >> 1
//         }
//         res = append(res,temp)
//     }
//     return res
// }

func subsets(nums []int) [][]int {
    subsets := [][]int{}
    backtracking(0,nums,[]int{},&subsets)
    return subsets
}
func backtracking(index int , nums []int, subs []int, subsets *[][]int){
    b := make([]int, len(subs))
    copy(b, subs)
    *subsets = append(*subsets,b)
    for i := index ; i < len(nums) ; i++{
        subs = append(subs,nums[i])
        backtracking(i+1,nums,subs,subsets)
        subs = subs[:len(subs)-1]
    }
}