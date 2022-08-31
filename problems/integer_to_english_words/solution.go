var s = []string{"","One","Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen","Fourteen","Fifteen","Sixteen","Seventeen", "Eighteen","Nineteen"}
var m = []string{"","Ten", "Twenty","Thirty","Forty","Fifty","Sixty","Seventy","Eighty","Ninety"}
var l = []string{"", "Thousand","Million","Billion"}
func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    result := []string{}
    count := 0
    for num > 0 {
        if num % 1000 != 0 {
            
            result = append(getString(num%1000), append([]string{l[count]}, result...)...)
        }
        num /= 1000
        count++
    }
    a := ""
    for _,i := range result {
        if i != " " && i != "" {
            a += i + " "
        }
    }
    return strings.Trim(a," ")
}
func getString( n int) []string{
    if n == 0 {
        return []string{""}
    } else if n < 20 {
        return []string{s[n]}
    } else if n < 100 {
        return append([]string{m[n/10]}, getString(n%10)...)
    } else {
        return append([]string{s[n/100] ,"Hundred"} ,getString(n % 100)...)
    }
}