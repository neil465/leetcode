func calculateMinimumHP(dungeon [][]int) int {
    v, k := re(dungeon, 0, 0, 0, 0, []int{}, map[[2]int]int{})
    fmt.Println(v, k)
   return -1 * v + 1
}
func re(d [][]int, curr, curc int, curHealth, minHealth int, arr []int, dp map[[2]int]int) (int, []int) {
    
    if curr >= len(d) || curc >= len(d[0]){
        return math.MinInt32, arr
    }
    if _, ok := dp[[2]int{curr, curc}]; ok {
        // fmt.Println("hi")
        return dp[[2]int{curr, curc}], []int{}
    }
    curHealth += d[curr][curc]
    arr = append(arr, d[curr][curc])
    if curHealth < minHealth {
        
        minHealth = curHealth
    }
    if curr == len(d) - 1 && curc == len(d[0]) -1 {
        // fmt.Println(arr)
        return minHealth, arr
    }
    v1, tarr1 := re(d, curr + 1, curc, 0, 0, arr, dp  )
    v2, tarr2 := re(d, curr, curc + 1, 0, 0, arr, dp)
    if v1 > v2 {
        if v1 + curHealth < minHealth {
            dp[[2]int{curr, curc}] = v1 + curHealth
            return v1 + curHealth,tarr1 
        }
        dp[[2]int{curr, curc}] = minHealth
        return minHealth, tarr1
        
    } 
    if v2 + curHealth < minHealth {
        dp[[2]int{curr, curc}] =v2 + curHealth
        return v2 + curHealth,tarr2 
    }
    dp[[2]int{curr, curc}] = minHealth
    return minHealth, tarr2

}