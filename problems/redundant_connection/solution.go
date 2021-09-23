type unionFind struct {
	parent []int
	n      int
}

func (this *unionFind) Find(x int) int {
	if x != this.parent[x] {
		this.parent[x] = this.Find(this.parent[x])
		return this.parent[x]
	}

	return x
}

func (this *unionFind) Union(x int, y int) {
	temp1 := this.Find(x)
	temp2 := this.Find(y)

	if temp1 != temp2 {
		this.n--
		this.parent[temp1] = temp2
	}
}


func findRedundantConnection(edges [][]int) []int {
    a := make([]int,len(edges))
    u := &unionFind{a,len(a)}
    for i := 0 ; i < len(edges) ; i++{
        u.parent[i]=i
    }
    for _,i := range edges{
        old := u.n
        u.Union(i[0]-1,i[1]-1)
        if old == u.n{
            return i
        }
    }
    return edges[0]
}