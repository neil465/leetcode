func accountBalanceAfterPurchase(purchaseAmount int) int {
    if purchaseAmount % 10 == 0 {
        return 100 - purchaseAmount
    } 
    b := 0
    if purchaseAmount % 10 > 4 {
        b = 1
    }
    return 100 -( purchaseAmount / 10 + b) * 10
}