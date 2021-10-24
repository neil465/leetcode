func isAnagram(s string, t string) bool {
    l := [26]int{}
    for _,i := range s{
        l[int(i-'a')]++
    }
    for _,i := range t{
        l[int(i-'a')]--
        if l[int(i-'a')] < 0{return false}
    }
    for _,i := range l{
        if i != 0{
            return false
        }
    }
    return true
}