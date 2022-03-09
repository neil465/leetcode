func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    
    for _,i := range strs{
        m[SortString(i)] = append(m[SortString(i)],i)
    }
    result := [][]string{}
    
    for _,i := range m {
        result = append(result,i)
    }
    return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}
