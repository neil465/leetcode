import "math"
func reverse(x int) int {
   

    res := 0
    flag := 1
    if x < 0 {
        flag = -1
        x *= -1     
    }
    count :=int(math.Pow(float64(10),float64((int(math.Log10(float64(x)))))))
    
    for x > 0{
        res += (x%10)*count
        count /= 10
        x /= 10
    }
    if res*flag >= 2147483647 || res*flag <= -2147483648{return 0}

    return res*flag
}	