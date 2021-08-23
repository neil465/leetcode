type HitCounter struct {a []int}


/** Initialize your data structure here. */
func Constructor() HitCounter {return HitCounter{}}


/** Record a hit.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int)  {this.a = append(this.a,timestamp)}


/** Return the number of hits in the past 5 minutes.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
    res := 0
    for _,i := range this.a{
        if i > timestamp-300 && i <= timestamp{res++}
    }
    return res
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */