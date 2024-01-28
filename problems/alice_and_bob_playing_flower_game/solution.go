func flowerGame(n int, m int) int64 {
    // // dp := make([][]int64, n)
    // // for i := range dp {
    // //     dp[i] = make([]int64, m)
    // // }
    // ma := map[[2]int]bool{}
    // for curn := 0 ; curn < n; curn ++ {
    //     for curm := 0 ; curm < m; curm ++ {
    //         // add := int64(0) 
    //         if (curn + curm) % 2 != 0 {
    //             // fmt.Println(curn, curm)
    //                 ma[[2]int{curn, curm}] = true
    //         }
    //         // if curn - 1 >= 0 {
    //         //     dp[curn][curm] += dp[curn - 1][curm] + add
    //         // }
    //         // if curm - 1 >= 0 {
    //         //     dp[curn][curm] += dp[curn][curm - 1] + add
    //         // }
    //     }
    // }
    // // return dp[n - 1][m - 1]
    // return int64(len(ma))
    return int64((n * m) / 2)
}