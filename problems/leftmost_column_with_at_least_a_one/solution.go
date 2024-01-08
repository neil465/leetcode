/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get func(int, int) int
 *     Dimensions func() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    rl, cl := binaryMatrix.Dimensions()[0], binaryMatrix.Dimensions()[1]
    minC := math.MaxInt32
    r := 0
    for i := 0 ; i < rl; i++ {
        
        ind := sort.Search(cl , func(colInd int) bool { 
            return binaryMatrix.Get(i, colInd) >= 1 
        })
        if ind < minC && ind < cl && binaryMatrix.Get(i, ind) == 1{
            minC = ind
            r = i
        }
    }
    if minC == math.MaxInt32 {
        return -1
    }
    if binaryMatrix.Get(r, minC) == 0 {
        return - 1
    }
    return minC
}