/**
 * Definition for a category handler.
 * type CategoryHandler interface {
 *  HaveSameCategory(int, int) bool
 * }
 */
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
func numberOfCategories(n int, categoryHandler CategoryHandler) int {
    a := make([]int, n)

    for i := 0 ; i< n ; i ++ {
        a[i] = i
    }
    uf := &unionFind{n, a}

    for i := 0 ; i< n ; i ++ {
        for j := 0 ; j< n ; j ++ {
            if categoryHandler.HaveSameCategory(i, j) {
                uf.Union(i, j)
            }
        }
    }
    return uf.count
}