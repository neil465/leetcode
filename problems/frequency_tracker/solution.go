type FrequencyTracker struct {
    m map[int]int    
    f map[int]int
}


func Constructor() FrequencyTracker {
    return FrequencyTracker{map[int]int{}, map[int]int{}}
}


func (this *FrequencyTracker) Add(number int)  {
    this.f[this.m[number]]--
    this.f[this.m[number] + 1]++
    this.m[number]++
}


func (this *FrequencyTracker) DeleteOne(number int)  {
    if this.m[number] > 0 {
        this.f[this.m[number]]--
        this.f[this.m[number] - 1] ++
        this.m[number]--
    }
    
}


func (this *FrequencyTracker) HasFrequency(frequency int) bool {

    if this.f[frequency] > 0 {
        return true
    }
    return false
}


/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */