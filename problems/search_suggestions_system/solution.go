func suggestedProducts(products []string, searchWord string) [][]string {
    result := [][]string{}
    for i := 1; i <= len(searchWord); i++{
        temp := []string{}
        for _,j := range products {
            if strings.HasPrefix(j, searchWord[:i]){
                temp = append(temp, j)
            }
        }
        sort.Strings(temp)
        if len(temp) > 3 {
            temp = temp[:3]
        }
        result = append(result,temp)
    }
    return result
}