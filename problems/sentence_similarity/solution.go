func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
    m := map[string][]string{}
    for _, i := range similarPairs {

        m[i[0]] = append(m[i[0]], i[1])
        m[i[1]] = append(m[i[1]], i[0])
    }
    if len(sentence1) != len(sentence2) {
        return false
    }
    for i := range sentence1 {
        f := true
        if sentence1[i] == sentence2[i] {
            continue 
        }
        for _, j := range m[sentence1[i]] {
            if j == sentence2[i]  {
                f = false
            }
        }
        if f {
            return false
        }
        
    }
    return true
}