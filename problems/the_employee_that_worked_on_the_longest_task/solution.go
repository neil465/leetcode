func hardestWorker(n int, logs [][]int) int {
    start, max, maxid := 0, 0, 0
    
    for _, i := range logs {
        if i[1] - start > max || (i[1] - start == max && i[0] < maxid){
            fmt.Println(i[1], start, max)
            max, maxid = i[1] - start, i[0]
        } 
        start = i[1] 
    }
    return maxid
}