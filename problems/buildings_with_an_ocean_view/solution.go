func findBuildings(heights []int) []int {
    res := []int{}
    for i:= 0 ; i < len(heights)-1; i++{
        flag := true
        for j := i+1 ; j < len(heights) ; j++{
            if heights[j] >= heights[i]{
                flag = false
                break
            }
        }
        if flag{
            res =append(res,i)
        }
    }
    res = append(res,len(heights)-1)
    return res
}