func addSpaces(s string, spaces []int) string {
    arr := make([]string, len(spaces)+1)
    count := 0 
    for i, j := range spaces{
        arr[i] = s[count:j]
        count = j
    }
    arr[len(arr) - 1] = s[count:]

    return strings.Join(arr, " ")
}