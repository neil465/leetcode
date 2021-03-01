type MyStack struct {
	foo []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.foo = append(this.foo, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.foo[len(this.foo)-1:]
	this.foo = this.foo[:len(this.foo)-1]
	return res[0]
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.foo[len(this.foo)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.foo)==0
}