func findPattern(stream InfiniteStream, pattern []int) int {
    // Preprocess the pattern to create the LPS (Longest Proper Prefix which is also Suffix) array.
    lps := preprocessPattern(pattern)

    j := 0 // Index for pattern
    cur := []int{stream.Next()} // Current window of stream data

    // Iterate through the stream
    for i := 0; ; {
        // Extend the current window if necessary
        if len(cur) <= i {
            cur = append(cur, stream.Next())
        }

        // Pattern matching logic
        if pattern[j] == cur[i] {
            j++ // Move to the next character in the pattern
            if j == len(pattern) {
                // Pattern found
                return i - len(pattern) + 1
            }
            i++ // Move to the next character in the stream
        } else if j > 0 {
            j = lps[j-1] // Use LPS array to skip characters
        } else {
            i++ // No match, move to the next character in the stream
        }
    }
}


func preprocessPattern(pattern []int) []int {
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