func vowelStrings(words []string, left int, right int) int {
    vowels := map[byte]bool{'a' : true, 'e' : true, 'i' : true, 'o' : true, 'u' : true}
    sum := 0
    for i := left; i <= right; i++{
        if vowels[words[i][0]] && vowels[words[i][len(words[i]) - 1]] {
            sum++
        }
    }

    return sum
}