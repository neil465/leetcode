func minOperations(logs []string) int {
    s := []int{}
    for _,i := range logs{
        if i == "../"{
            if len(s) > 0 {
                s = s[:len(s)-1]
            }
            
        }else if i!= "./"{
            s = append(s,0)
        }
    }
    return len(s)
}