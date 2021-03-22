
import "sort"

func getMinimumDifference(root *TreeNode) int {
	stack := []*TreeNode{root}
	arr := []int{}
	for len(stack) > 0 {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if value != nil {
			arr = append(arr, value.Val)
			stack = append(stack, value.Left, value.Right)
		}
	}

	sort.Ints(arr)
	minimum := abs(arr[0], arr[1])
	for i := 0; i < len(arr)-1; i ++ {
		if abs(arr[i], arr[i+1])<minimum {
			minimum = abs(arr[i], arr[i+1])
		}
	}
	return minimum
}
func abs(val int, val2 int) int {
	if val-val2 < 0 {
		return -(val - val2)
	}
	return val - val2
}