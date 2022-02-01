type unionFind struct {
	count   int
	parents []int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
		this.count--
	}
}


func earliestAcq(logs [][]int, n int) int {
    sort.Slice(logs,func(i,j int)bool{
        return logs[i][0] < logs[j][0]
    })
    arr := make([]int,n)
    union := &unionFind{count : n,parents : arr}
    for i := 0 ; i < n ; i++{
        union.parents[i] = i
    }
    for _,i := range logs{
        union.Union(i[1],i[2])
        if union.count == 1 {
            return i[0]
        }
    }
    return -1
}