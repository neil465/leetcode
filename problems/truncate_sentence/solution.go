import "strings"

func truncateSentence(s string, k int) string {
	arr :=strings.Split(s," ")
	result := ""
	for i := 0; i < len(arr); i++ {
		if i<=k-1 {
			result += arr[i]+" "
		}
	}
	return strings.TrimRight(result," ")
}