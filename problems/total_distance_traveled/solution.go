func distanceTraveled(mainTank int, additionalTank int) int {
    a := 0
    for mainTank > 0 {
        if mainTank < 5 {
            a += mainTank * 10 
            mainTank = 0
            break
        }
        a += 5 * 10 
        mainTank -= 5
        if additionalTank > 0 {
            mainTank++
            additionalTank --
        }
    }
    return a
}