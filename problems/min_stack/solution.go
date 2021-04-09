
type MinStack struct {
	foo  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.foo = append(this.foo,val)
	
}

func (this *MinStack) Pop() {
	this.foo = this.foo[:len(this.foo)-1]
}

func (this *MinStack) Top() int {
		return this.foo[len(this.foo)-1]
}

func (this *MinStack) GetMin() int {
	minumum := this.foo[0]
	for _, i2 := range this.foo {
		if i2<minumum {
			minumum= i2
		}
	}
	return minumum
}