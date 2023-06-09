/**
 * Definition for a street.
 * type Street interface {
 *     OpenDoor()
 *     CloseDoor()
 *     IsDoorOpen() bool
 *     MoveRight()
 *     MoveLeft()
 * }
 */
func houseCount(street Street, k int) int {
    for i := 0 ; i < k ; i++ {
        street.CloseDoor()
        street.MoveRight()
    }
    a := 0
    for street.IsDoorOpen() == false {
        street.OpenDoor()
        street.MoveRight()
        a++
    }
    return a
}