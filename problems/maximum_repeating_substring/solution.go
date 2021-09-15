func maxRepeating(sequence string, word string) int {
    temp := word
    count := 1
    for len(temp) <= len(sequence){
        flag := true
        for i := 0 ; i < len(sequence)-len(temp)+1 ; i++{
            if temp == sequence[i:i+len(temp)]{
                flag = false
                break
            }
        }
        if flag {
            return count-1
        }
        temp += word
        count++
    }
    return count-1
}