func killProcess(pid []int, ppid []int, kill int) []int {
    res := []int{kill}
    toBeKilled := []int{kill}
    m := make(map[int][]int)
    for i := 0 ; i < len(ppid) ; i ++{
        temp := m[ppid[i]]
        temp = append(temp,pid[i])
        m[ppid[i]] = temp
    }
    for len(toBeKilled) > 0{
        arr := []int{}
        for _,i := range toBeKilled{
            if len(m[i]) > 0{
                res = append(res,m[i]...)
                arr = append(arr,m[i]...)
            }
        }
        toBeKilled = arr
    }
    return res
    
}