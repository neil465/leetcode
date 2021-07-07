

func isMonotonic(A []int) bool {
    flag := -1
    for i:= 0 ; i < len(A)-1 ; i++{
        if flag == -1 && A[i] > A[i+1]{
            flag = 1
        }
        if flag == -1 && A[i] < A[i+1]{
            flag = 0
        }
        if flag == 0 && A[i] > A[i+1]{
            return false
        }
        if flag == 1 && A[i] < A[i+1]{
            return false
        }
    }
	return true
}
