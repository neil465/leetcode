type person struct {
    name string
    height int
}

func sortPeople(names []string, heights []int) []string {
    a := []person{}
    for i := range names {
        a = append(a,person{names[i], heights[i]})
    }
    sort.Slice(a , func(i,j int) bool {
        return a[i].height > a[j].height
    })
    res := []string{}
    for _, i := range a {
        res = append(res, i.name)
    }
    return res
}