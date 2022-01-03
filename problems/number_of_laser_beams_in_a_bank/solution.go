func numberOfBeams(bank []string) int {
    a := fmt.Sprint(len(bank[0]))
    
    s := fmt.Sprintf("%0"+a+"d", 0)
    count := 0
    rowbefore := 0
    for _,i := range bank{
        if i == s{
            continue
        }
        row := 0
        for _,j := range i{
            if j == '1'{row++}
        }
        fmt.Println(row,rowbefore,s,i)
        count += row*rowbefore
        rowbefore = row
    }
    return count
}