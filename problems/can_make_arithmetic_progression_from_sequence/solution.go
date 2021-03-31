
import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[0]-arr[1]
	for i := 1; i < len(arr); i++{
		if arr[i-1]-arr[i] != diff {
			return false
		}
	}
	return true
}