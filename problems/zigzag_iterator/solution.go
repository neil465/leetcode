type ZigzagIterator struct {
    a []int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    a := []int{}
    z := int(math.Min(float64(len(v1)), float64(len(v2))))
    for i := 0 ; i < z; i++ {

        a = append(a, v1[0], v2[0])
        v1 = v1[1:]
        v2 = v2[1:]
    }
    a = append(a, v1...)
    a = append(a, v2...)
    return &ZigzagIterator{a}
}

func (this *ZigzagIterator) next() int {
    v := this.a[0]
    this.a = this.a[1:]
    return v
}

func (this *ZigzagIterator) hasNext() bool {
    return len(this.a) > 0
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */