func countMatchingSubarrays(nums []int, pattern []int) int {
    // Preprocess the pattern to create the LPS (Longest Proper Prefix which is also Suffix) array.
    lps := preprocessPattern(pattern)

    j := 0 // Index for pattern

    i := 0
    sum := 0
    for i < len(nums)- 1{

        // Pattern matching logic

        v := 0

        if nums[i] < nums[i + 1] {
            v = 1
        } else if nums[i] > nums[i + 1] {
            v = -1
        }
        if pattern[j] == v {
            j++ // Move to the next character in the pattern
            if j == len(pattern) {
                // Pattern found
                sum ++
                j = lps[j - 1]
            }
            i++ // Move to the next character in the stream
        } else if j > 0 {
            j = lps[j-1] // Use LPS array to skip characters
        } else {
            i++ // No match, move to the next character in the stream
        }
    }
    return sum
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