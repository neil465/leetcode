func subdomainVisits(cpdomains []string) []string {
    m := map[string]int{}
    for _, countPair := range cpdomains {
        value := strings.Fields(countPair)
        num, _ := strconv.Atoi(value[0])
        subDomains := strings.Split(value[1], ".")

        for i := range subDomains {
            m[strings.Join(subDomains[i:], ".")] += num
        }
    }
    a := []string{}
    for d, v := range m {
        a = append(a, strconv.Itoa(v) + " " + d)
    }
    return a
}