type name struct {
	K int
	V int
}

func topKFrequent(nums []int, k int) []int {
	maper := make(map[int]int)

	topk := []int{}
	for _, i := range nums {
		maper[i]++
	}
	a := make([]name, len(maper))
	count := 0
	for i, j := range maper {
		a[count].K = i
		a[count].V = j
		count++
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].V == a[j].V {
			return a[i].K < a[j].K
		}
		return a[i].V > a[j].V
	})


	for i := 0; i < k; i++ {
		topk = append(topk, a[i].K)
	}

	return topk
}