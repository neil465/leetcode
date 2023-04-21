type Row struct {
    value, id int
}

func kWeakestRows(mat [][]int, k int) []int {
    a := []Row{}
    for index, i := range mat {
        s := 0
        for _, j := range i {
            s += j
        }
        a = append(a, Row{s, index})
    }
    sort.Slice(a, func(i,j int) bool {
        if a[i].value == a[j].value {
            return a[i].id < a[j].id
        }
        return a[i].value < a[j].value
    })
    res := []int{}
    for i := 0; i< k ; i++ {
        res = append(res, a[i].id)
    }
    return res
}