func numUniqueEmails(emails []string) int {
   
    maper := make(map[string]int)
    for _,i := range emails{
        d := strings.Split(i,"@")
        domain := d[1]
        name := d[0]
        name=strings.ReplaceAll(name,".","")
        name = strings.Split(name,"+")[0]
        email := name +"@" + domain
        maper[email] = 0
    }
    return len(maper)
    
}