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

func makeConnected(n int, connections [][]int) int {
    extra := 0 
    parents := make([]int, n)
    for i := 0 ; i < n ; i++ {
        parents[i] = i
    }
    uf := &unionFind{n, parents}
    for _, i := range connections {
        if uf.Find(i[0]) == uf.Find(i[1]) {
            extra ++
        } else {
            uf.Union(i[0], i[1])
        }
    }
    if uf.count -1 <= extra {
        return uf.count - 1
    }
    return -1

}