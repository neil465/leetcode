func beautifulIndices(s string, a string, b string, k int) []int {
     // Preprocess the pattern to create the LPS (Longest Proper Prefix which is also Suffix) array.
    alps := preprocessPattern([]byte(a))
    blps := preprocessPattern([]byte(b))

    j := 0 // Index for pattern

    i := 0
    inds := []int{} 
    for i < len(s){

        // Pattern matching logic

        
        if a[j] == s[i] {
            j++ // Move to the next character in the pattern
            if j == len(a) {
                // Pattern found
                inds = append(inds, i - len(a) + 1)
                j = alps[j - 1]
            }
            i++ // Move to the next character in the stream
        } else if j > 0 {
            j = alps[j-1] // Use LPS array to skip characters
        } else {
            i++ // No match, move to the next character in the stream
        }
    }

    j = 0 // Index for pattern

    i = 0

    inds2 := []int{}

    for i < len(s){

        // Pattern matching logic

        
        if b[j] == s[i] {
            j++ // Move to the next character in the pattern
            if j == len(b) {
                // Pattern found
                inds2 = append(inds2, i - len(b) + 1)
                j = blps[j - 1]
            }
            i++ // Move to the next character in the stream
        } else if j > 0 {
            j = blps[j-1] // Use LPS array to skip characters
        } else {
            i++ // No match, move to the next character in the stream
        }
    }

    res := []int{}

    sort.Ints(inds2)
    for i := range inds {
        
        if binarySearchNearby(inds2, inds[i], k) {
            res = append(res, inds[i])
        }
    }


    return res
}

func binarySearchNearby(arr []int, target, k int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if math.Abs(float64(arr[mid]-target)) <= float64(k) {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func preprocessPattern(pattern []byte) []int {
    lps := make([]int, len(pattern))
    for i, j := 1, 0; i < len(pattern); {
        if pattern[i] == pattern[j] {
            j++
            lps[i] = j
            i++
        } else if j > 0 {
            j = lps[j-1]
        } else {
            lps[i] = 0
            i++
        }
    }
    return lps
}