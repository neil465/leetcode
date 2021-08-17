/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    getVal := 0
    for reader.get(getVal) != 2147483647{
        if reader.get(getVal) == target{
            return getVal

        }
        getVal++
    }
    return -1

}