type key struct{
    k,v int
}

func findRelativeRanks(score []int) []string {
    arr := []key{}
    for i,j := range score{
        arr = append(arr,key{i,j})
    }
    sort.Slice(arr,func(i,j int)bool{
        return arr[i].v >arr[j].v
    })
    fmt.Println(arr)
    res := make([]string , len(arr))
    for i := 0 ; i < len(arr) ; i ++{
        if i == 0{
            res[arr[i].k] = "Gold Medal"
        }else if len(arr) > 1 && i == 1{
            res[arr[i].k] = "Silver Medal"
        }else if len(arr) > 2 && i == 2{
            res[arr[i].k] = "Bronze Medal"
        }else{
            res[arr[i].k] = fmt.Sprint(i+1)
        }
        
    }
    return res
}