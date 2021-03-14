func replaceElements(arr []int) []int {
	res := []int{}
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			res = append(res, -1)
		}else{
			max := arr[i+1]
			for j := i+2; j < len(arr); j++ {
				if arr[j]>max {
					max = arr[j]
				}
			}
			res = append(res,max)
		}
		
	}
	return res
}