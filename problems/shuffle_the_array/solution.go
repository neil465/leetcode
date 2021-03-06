import "fmt"
func shuffle(nums []int, n int) []int {
	result := []int{}
	yarr := nums[n:]
	xarr := nums[:n]
	for i := 0; i < len(xarr); i++ {
		result = append(result,xarr[i])
		result = append(result,yarr[i])
	}
    
	return result

}