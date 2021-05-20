func uncommonFromSentences(s1 string, s2 string) []string {
    res := []string{}
                                                           
    mapy := make(map[string]int)
    s3 := s1 + " " + s2
    s32 := strings.Split(s3," ")
    for _ , i2 := range s32{
        mapy[string(i2)]++
    }
    for k,_ := range mapy {
        if mapy[string(k)] < 2{
            res = append(res ,string(k))
        }
    }
    return res
}