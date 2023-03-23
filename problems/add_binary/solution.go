func addBinary(ta string, tb string) string {
    fmt.Println(ta, tb)
    if len(ta) == len(tb) && len(ta) == 1 {
        if string(ta[0]) == "0" && string(tb[0]) == "0" {
            return "0"
        } else if string(ta[0]) == "0" && string(tb[0])== "1"|| string(ta[0]) == "1" && string(tb[0]) == "0"{
            return "1"
        } else {
            return "10"
        }
    }

    if len(ta)< len(tb) {
        ta = strings.Repeat("0", len(tb) - len(ta)) + ta
    } else {
        tb = strings.Repeat("0", len(ta) - len(tb)) + tb
    }
    ta = "0" + ta
    tb = "0" + tb
    i:= len(ta) - 1
    carry := 0
    sum := ""
    fmt.Println(ta,tb)
    for i >= 0 {

        val := int(ta[i]) - 48 + int(tb[i]) - 48 + carry
        if val >= 2 {
            carry = 1
            if val == 3 {
                sum = "1" + sum
            } else {
                sum = "0" + sum
            }
        } else {
            carry = 0
            fmt.Println("s1" ,sum)
            sum = string(val + 48) + sum
            fmt.Println("s2", sum)
        }
        fmt.Println(sum, "sum")
        i--

        
        
    }
    sum = strings.TrimLeft(sum,"0")
    return sum
}