type ParkingSystem struct {
    arr [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{[3]int{big,medium,small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if carType == 3{
        if this.arr[2] > 0{
            this.arr[2] = this.arr[2]-1
            return true
        }
    }
    if carType == 1{
        if this.arr[0] > 0{
            this.arr[0] = this.arr[0]-1
            return true
        }
    }
    if carType == 2{
        if this.arr[1] > 0{
            this.arr[1] = this.arr[1]-1
            return true
        }
    }
    return false
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */