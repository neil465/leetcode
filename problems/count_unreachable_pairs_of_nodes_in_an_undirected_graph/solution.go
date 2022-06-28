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

func countPairs(n int, edges [][]int) int64 {
    c:= 0
    u := &unionFind{n, make([]int, n)}
    for i := 0 ; i < n ; i++{
        u.parents[i] = i
    }
    for i := range edges {
        u.Union(edges[i][0], edges[i][1])
    }
    m := map[int]int{}
    for _,i := range u.parents {
        m[u.Find(i)] ++
    }
    a := []int{}
    for _, i := range m {
        a = append(a, i )
    }
    for i := 0 ; i < len(a) ; i++{
        for j := i + 1 ; j < len(a) ; j++{

            c += a[i] * a[j]
        }
    }
    return int64(c)
}
