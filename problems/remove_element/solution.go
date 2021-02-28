func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	//fmt.Println(nums)
	flag := true
	counter := 0
	n := len(nums)
	for i := 0; i < n; i++ {

		if nums[counter] == val {
			flag = false
			//fmt.Println("foo", nums[counter], counter)
			if counter == len(nums)-1 {
				//fmt.Println("inside", nums[counter])
				nums = nums[:len(nums)-1]
			} else {
				nums = append(nums[:counter], nums[counter+1:]...)
				//fmt.Println("foo2", "nums", nums)
			}

		}
		if flag {
			counter++
		}

	}
	//fmt.Println(nums)
	return len(nums)
}