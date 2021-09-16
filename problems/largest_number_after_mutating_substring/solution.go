func maximumNumber(num string, change []int) string {
    flag := false
    arr := strings.Split(num, "")
    for i := range arr {
        temp, _ := strconv.Atoi(arr[i])
        if change[temp] > temp {
            arr[i] = strconv.Itoa(change[temp])
            flag = true
        } else if flag && change[temp] == temp {
            continue
        } else if flag{
            break
        }
    }

    return strings.Join(arr,"")
}