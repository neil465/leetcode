func nextGreatestLetter(letters []byte, target byte) byte {
    if letters[len(letters) - 1] <= target {
        return letters[0]
    }
    l, r := 0, len(letters) - 1
    mid := (l + r) / 2
    for l <= r {
        if letters[mid] <= target {
            l = mid + 1
        } 
        if letters[mid] > target {
            r = mid - 1
        }
        mid = (l + r) / 2
    }
    return letters[l]
}