func frequencySort(s string) string {
    freq := map[rune]int{}

    for _, i := range s{
        freq[i] ++
    }

    runes := []rune(s)

    sort.Slice(runes, func(i, j int) bool {
        if freq[runes[i]] == freq[runes[j]] {
            return runes[i] < runes[j]
        }
        return freq[runes[i]] > freq[runes[j]]
    })


    return string(runes)
}