import "strconv"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
    first,second,target := "","",""
    for _,i := range firstWord{
        first+= fmt.Sprint(int(rune(i))-97)
    }
    for _,i := range secondWord{
        second += fmt.Sprint(int(rune(i))-97)
    }
    for _,i := range targetWord{
        target += fmt.Sprint(int(rune(i))-97)
    }
    iirst , _ := strconv.Atoi(first)
    iecond , _ := strconv.Atoi(second)
    sum := iirst + iecond
    itarget , _ := strconv.Atoi(target)
    if sum == itarget{
        return true
    }
    return false
}