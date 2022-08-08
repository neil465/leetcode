func intToRoman(num int) string {
    m := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
    v := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
    s := ""
    for i := 0 ; i < len(v);i++ {
        if num >= v[i ]{
            num -= v[i]
            s+= m[i]
            i--
        }
    }
    return s

    
}