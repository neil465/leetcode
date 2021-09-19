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
func countComponents(n int, edges [][]int) int {
    a := []int{}
    for i := 0 ; i<n;i++{
        a = append(a,i)
    }
    u := &unionFind{a,n}
    for _,i := range edges{
        u.Union(i[0],i[1])
    }
    return u.count
}