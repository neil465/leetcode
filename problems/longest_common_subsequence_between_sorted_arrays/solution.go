func longestCommomSubsequence(arrays [][]int) []int {
    arr := make(map[int]int)
    for _,i := range arrays{
        for _,j := range i{
            arr[j]++
        }

    }
    b := []int{}
    for k,v := range arr{
        if v == len(arrays){
            b =append(b,k)
        }
    }
    sort.Ints(b)
    return b
}