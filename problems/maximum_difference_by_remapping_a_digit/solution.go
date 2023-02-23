func minMaxDifference(num int) int {
    a := strconv.Itoa(num)

    var max int
    for i := range a {
        if string(a[i]) != "9" || i == len(a) - 1  {
            max, _ = strconv.Atoi(strings.Replace(a, string(a[i]), "9", 3000000))
            break
        }
    }
    min, _ := strconv.Atoi(strings.Replace(a, string(a[0]), "0", 3000000))

    return max - min
}