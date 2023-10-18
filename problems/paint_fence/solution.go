var dp = map[[2]int]int{}

func numWays(n int, k int) int {
    dp = map[[2]int]int{}
    return find(n, k, 0, []int{})
}
func find(n, k, consecCount int, curarr []int) int {
    if n== 0 {
        return 1
    }
    
    if _, ok := dp[[2]int{n, consecCount}]; ok {

        return  dp[[2]int{n, consecCount,}]
    } 

        

     
    
    val := 0
    for i := 0; i < k; i ++ {
        if consecCount == 2 {
            if curarr[len(curarr) - 1] != i {
                val += find(n - 1, k, 1 , append(curarr, i))    
            }
        } else {
            if len(curarr) > 0 && curarr[len(curarr) - 1] == i {
                val += find(n - 1, k, consecCount + 1, append(curarr, i))   
            } else {

                    val += find(n - 1, k, 1, append(curarr, i))   

            }
        }
    }


        dp[[2]int{n, consecCount}] = val



   
    return val

}