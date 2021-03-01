
type MyQueue struct {
	 foo []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.foo = append(this.foo,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.foo[0]
	this.foo  = this.foo[1:]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.foo[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.foo) == 0
}