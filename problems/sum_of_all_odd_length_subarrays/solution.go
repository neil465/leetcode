func sumOddLengthSubarrays(arr []int) int {
	res := 0
	for i := 1; i < len(arr)+1; i += 2 {

		for i2, _ := range arr {
			if len(arr)>=i2+i {
				split := arr[i2 : i2+i]
                fmt.Println(split)
				for _, i3 := range split { //adding each part of the subslice to res
					res += i3
				}
                
			}else {
				continue
			}
			
		}
	}
	return res
}
