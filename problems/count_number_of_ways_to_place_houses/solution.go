func countHousePlacements(n int) int {
    a1,a2 := 0 ,1
    sum := a1 + a2
    for i := 1; i < n + 2 ; i++{
        a1, a2 = a2, sum

        sum = (a1+a2) %1000000007
    }
    return (a2 ) * (a2) %1000000007
} 
