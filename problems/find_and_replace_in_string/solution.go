type group struct {
    index int
    source, target string
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    groups := []group{}
    for i := range indices {
        groups = append(groups, group{indices[i], sources[i], targets[i]} )
    }
    sort.Slice(groups, func(i, j int) bool {
        return groups[i].index < groups[j].index
    })
    changeIndex := 0 
    for index := range groups {
        if index == 1 {
            fmt.Println(groups[index].index + changeIndex)
        }
        if groups[index].index + changeIndex + len(groups[index].source) <= len(s) && s[groups[index].index + changeIndex: groups[index].index + changeIndex + len(groups[index].source)] == groups[index].source {
            change := len(groups[index].target) - len(groups[index].source)
            s = s[:groups[index].index + changeIndex] + groups[index].target + s[groups[index].index + changeIndex + len(groups[index].source):]
            changeIndex += change
        }
    }
    return s
}