type temp struct {
	arr []int
	k bool
}

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


func countCompleteComponents(n int, edges [][]int) int {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	uf := &unionFind{
		count:   n,
		parents: parents,
	}
	m := map[[2]int]bool{}
	for _, i := range edges {
		uf.Union(i[0], i[1])
		m[[2]int{i[0], i[1]}] = true
		m[[2]int{i[1], i[0]}] = true
	}
	count := uf.count
	arr := map[int]temp{}
	for i := 0; i < n; i++ {
		z := uf.Find(i)

		if arr[z].k {
			continue
		}

		k2 := false

		for _, a := range arr[z].arr {
			if !m[[2]int{a, i}] {
				k2 = true
			}
		}

		if k2 {
			arr[z] = temp{arr[z].arr, true}
			count --
		} else {
			arr[z] = temp{append(arr[z].arr, i), arr[z].k}
		}

	}
	return count
}
