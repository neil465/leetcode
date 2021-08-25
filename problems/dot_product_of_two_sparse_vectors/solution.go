
type SparseVector struct {
   a map[int]int
}

func Constructor(nums []int) SparseVector {
    m := make(map[int]int)
    for i,j := range nums{
        if j != 0 {m[i] = j}
    }
    return SparseVector{m}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    res := 0
    a := min(this.a,vec.a)
    iterate := a[0]
    non := a[1]
    fmt.Println(iterate,non)
    for k,i := range iterate{
        if v, ok := non[k]; ok {
       
            res += i*v
        }
    }
    return res
}
func min(i,j map[int]int) []map[int]int{
    if len(i) < len(j){return []map[int]int{i,j}}
    return []map[int]int{j,i}
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */