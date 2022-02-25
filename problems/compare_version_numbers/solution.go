func compareVersion(version1 string, version2 string) int {
    var v1, v2 = strings.Split(version1,"."),strings.Split(version2,".")
    
    for i := 0 ; i < int(math.Max(float64(len(v1)),float64(len(v2))))  ; i++{
        var a,b = 0,0
        if len(v1) > i { a,_ = strconv.Atoi(v1[i]) }
        if len(v2) > i { b,_ = strconv.Atoi(v2[i]) }
        if max(a,b) != 0 {return max(a,b)}
    }
    return 0
}

func max(i,j int) int {
    if i == j { return 0 }
    if i < j { return -1 }
    return 1
}