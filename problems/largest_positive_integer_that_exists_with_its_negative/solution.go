// in this solution, we are iterating through all the values in nums and checking if we have seen the additive inverse
// of a number, we then get the largest of all these numbers and return it

// walkthrough of the solution given the following input: [-1,2,-3,3]
// i = 0 : m = {-1:true} : max = -1
// i = 1 : m = {-1:true, 2:true} : max = -1
// i = 2 : m = {-1:true, 2:true, -3:true} : max = -1
// i = 3 : m = {-1:true, 2:true, -3:true, 3:true} : max = 3 // we found a valid solution, so we assign the absolute value of the current number to the max

// return the largest number that is the additive inverse of another number in the array : 3

func findMaxK(nums []int) int {
	// max is -1 because if there is no valid solution, we return -1
	m, max := map[int]bool{}, -1
	for _, i := range nums {
		// we add the current number to the map
		m[i] = true
		// if the additive inverse of the current number is in the map, and the absolute value of the current number is greater than the current max
		if m[-i] && abs(i) > max {
			// we assign the absolute value of the current number to the max
			max = abs(i)
		}
	}
	return max
}

func abs(i int) int { // helper function to get the absolute value of a number
	if i < 0 {
		return -i
	}
	return i
}
