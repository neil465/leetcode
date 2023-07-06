func largestMerge(word1 string, word2 string) string {
    index1, index2 := 0, 0 
    s := ""
    for index1 < len(word1) || index2 < len(word2) {

        
        
        if word1[index1:] > word2[index2:] {
            s += string(word1[index1])
            index1 ++
        } else {
            s += string(word2[index2])
            index2 ++
        }
        
        
    }
    return s
}