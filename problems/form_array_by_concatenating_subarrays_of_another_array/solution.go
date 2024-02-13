func canChoose(groups [][]int, nums []int) bool {
    lpss := make([][]int, len(groups))


   
    for index, _ := range lpss {
        lps := make([]int, len(groups[index]))
        j := 0 
        for i := 1 ; i < len(groups[index]); i ++ {
            for groups[index][j] != groups[index][i] && j > 0 {

                j = lps[j - 1]
            }

            if groups[index][j] == groups[index][i] {
                j ++
                lps[i] = j
            }
        }
        lpss[index] = lps
    }
    


    j := 0

    curGroup := 0


    // return false

    for i := 0 ; i < len(nums); i ++ {
        if curGroup >= len(groups) {
            return true
        }

        if groups[curGroup][j] == nums[i] {
            if j == len(groups[curGroup]) -1 {
                curGroup ++
                j = -1
            }
            
            j ++
        } else if j != 0 {
            j = lpss[curGroup][j - 1]
            i --
        }
       
    }
    return curGroup == len(groups)
}
