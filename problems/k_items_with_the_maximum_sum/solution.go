func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    if k <= numOnes {
        return k
    }
    if k <= numOnes + numZeros {
        return numOnes
    }
    if k <= numOnes + numZeros + numNegOnes {
        return numOnes - (k - (numOnes + numZeros))
    }
    return -1
}