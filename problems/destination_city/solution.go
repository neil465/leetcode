func destCity(paths [][]string) string {

    m := make(map[string]string)
    for _,i := range paths{
        m[i[0]] = i[1]
    }
    temp := paths[0][0]
    for true {
        if _, ok := m[temp]; ok {
            temp = m[temp]
        }else{
            break
        }
    }
    return temp
}