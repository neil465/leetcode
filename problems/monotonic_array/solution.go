import "sort"

func isMonotonic(A []int) bool {
	res := []int{}
	for i := 0; i < len(A); i++ {
        if i+1 < len(A) && A[i] > A[i+1] {
			res = append(res,1)
		}

		if i+1 < len(A) && A[i] < A[i+1] {
			res = append(res,0)
		}

	}
	sort.Ints(res)
	for i := 0; i < len(A); i++ {
		if i+1 < len(res) && res[i] != res[i+1]{
			return false
		}
	}
	return true
}
