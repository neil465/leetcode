// in this solution, we are iterating through all the nodes and assigning the value of the current node to its parent.
// This is because the value of a node is the same as the value of its parent unless there is a query for the current node.
// So if there is a query for the current node, we flip the current node's value. And if the node is true, or it is a one, we increment
// our one counter.

// walkthrough of the solution given the following input: 5, [1,2,5]

// i = 1 : parents = {1:true, 2:false, 3:false, 4:false, 5:false} : ones = 0 // have a query for 1, so we flip the value of 1
// i = 2 : parents = {1:true, 2:false, 3:false, 4:false, 5:false} : ones = 1 // have a query for 2, so we flip the value of 2, from 1's value to false
// i = 3 : parents = {1:true, 2:false, 3:true, 4:false, 5:false} : ones = 1 // we keep the value of 3 the same as 1 its parent, so 3 is true
// i = 4 : parents = {1:true, 2:false, 3:true, 4:false, 5:false} : ones = 2 // we keep the value of 4 the same as 2 its parent, so 4 is false
// i = 5 : parents = {1:true, 2:false, 3:true, 4:false, 5:true} : ones = 2 // have a query for 5, so we flip the value of 5, from 2's value to true

// return the number of ones : 3

func numberOfNodes(n int, queries []int) int {
	ones, q, parents := 0, map[int]int{}, map[int]bool{}
	// we are adding queries to a map so that we can retrieve them in constant time
	for _, i := range queries {
		q[i]++
	}
	for i := 1; i <= n; i++ {
		// we assign the value of the current node to its parent
		parents[i] = parents[i/2]
		// if there is a query for the current node, and the queries do not cancel each other out
		if q[i]%2 != 0 {
			// we flip the value of the current node
			parents[i] = !parents[i]
		}
		// if the value of the current node is 1, we increment the number of ones
		if parents[i] {
			ones++
		}
	}
	// we return the number of ones
	return ones
}
