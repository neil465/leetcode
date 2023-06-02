/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type pair struct {
    key, value int
}
type IntHeap []pair

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(pair))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var m = map[int]int{}

func findFrequentTreeSum(root *TreeNode) []int {
    m = map[int]int{}
    find(root)
    h := &IntHeap{}
    heap.Init(h)
    for k, v := range m {
        heap.Push(h, pair{k, v})
    }
    a := []int{}
    fmt.Println(m)
    lastVal := 0
    if h.Len() > 0 {
        z := heap.Pop(h).(pair)
        lastVal = z.value
        a = append(a, z.key)
    }
    
    for h.Len() > 0 {
        z := heap.Pop(h).(pair)
        if z.value == lastVal {
            a = append(a, z.key )
        } else {
            break
        }
        
    }
    return a

}
func find(root *TreeNode) int {
    if root == nil {
        return 0
    }
    v := find(root.Right) + find(root.Left) + root.Val
    m[v] ++
    return v
}