package main

// all we are doing in this solution is finding the minimum element that comes before pointer i, and check if this is the maximum difference.
func maximumDifference(nums []int) int {
	min, max := 0, -1 // min index and maximum dif

	for i := 1; i < len(nums); i++ {
		if nums[min] >= nums[i] { // if min is greater than or equal to the current index we have to move min up to the new minimum
			min = i
		}
		if nums[i] - nums[min] > max { // if the difference between the current index and the min is greater than the current max, we set the max to the new max
			max = nums[i] - nums[min]
		}
	}
	if max == 0 { // if the max is 0, that means that there was no difference between the numbers, so we return -1
		return -1
	}
	return max
}

/* Visual Explanation of Solution:

  max = -1         max = 0         max = 4         max = 4
7, 1, 5, 4		7, 1, 5, 4		7, 1, 5, 4    	7, 1, 5, 4
   |  |  	-> 	      |      ->    |  | 	->     |     |    ->  max = 4
  min i       		  i			  min i		      min    i
   					 min
*/
