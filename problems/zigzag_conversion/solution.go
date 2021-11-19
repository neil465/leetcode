func convert(s string, numRows int) string {
    if numRows == 1{return s}
    arr := make([][]string,numRows)
    p := 0
    add := 1
    for _,i := range s{
        arr[p] = append(arr[p],string(i))
        if p == 0 {add = 1}
        if p == numRows-1{add=-1}
        p += add
    }
    g := ""
    for _,i := range arr{for _,j := range i{g+= string(j)}}
    return g
}