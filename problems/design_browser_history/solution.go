type BrowserHistory struct {
    a []string
    index int
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{[]string{homepage}, 0}
}


func (this *BrowserHistory) Visit(url string)  {

    this.a = this.a[:this.index + 1]
    this.a = append(this.a, url)
    this.index++


}


func (this *BrowserHistory) Back(steps int) string {
    if steps > this.index  {
        this.index = 0
        return this.a[this.index]
    }

    this.index -= steps

    return this.a[this.index]

}


func (this *BrowserHistory) Forward(steps int) string {
    if steps + this.index >= len(this.a)  {
        this.index = len(this.a) - 1
        return this.a[this.index]
    } 
    this.index += steps
    return this.a[this.index]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */