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


func findCircleNum(isConnected [][]int) int {
    a := []int{}
    for i := 0; i < len(isConnected) ; i++{
        a = append(a,i)
    }
    
    b := &unionFind{a,len(isConnected)}
    for i,c := range isConnected{
        for j,k := range c{
            if k == 1{
                
                b.Union(i,j)    
            }
        }
    }
    return b.count
}