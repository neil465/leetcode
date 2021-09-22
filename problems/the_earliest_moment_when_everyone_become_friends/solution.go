type unionFind struct {
	parents []int
	count   int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}

	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x int, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx == ry {
		return
	}

	this.count--
	this.parents[rx] = ry
}

func earliestAcq(logs [][]int, n int) int {
    sort.Slice(logs,func(i,j int)bool{
        return logs[i][0] < logs[j][0]
    })
    a := make([]int,n)
    u := &unionFind{parents:a,count:n}   
    for i := 0 ; i < n ; i++{
        u.parents[i] = i
    }
    for _,i := range logs{
        u.Union(i[1],i[2])
        if u.count == 1{
            return i[0]
        }
    }
    return -1
    
}