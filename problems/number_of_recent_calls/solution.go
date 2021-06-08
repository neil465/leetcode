type RecentCounter struct {
    k []int
}


func Constructor() RecentCounter {
    return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
    res := 0
    this.k = append(this.k,t)
    for _,i := range this.k{
        if i >= t-3000{
            res++
        }
    }
    return res
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */