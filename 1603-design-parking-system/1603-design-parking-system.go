type ParkingSystem struct {
    bigSlot int
    midSlot int
    smallSlot int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        bigSlot : big,
        midSlot: medium,
        smallSlot: small,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if carType == 1 && this.bigSlot <= 0 {
        return false
    } else if carType == 1 {
        this.bigSlot--
        return true
    }
    if carType == 2 && this.midSlot <= 0 {
        return false
    } else if carType == 2 {
        this.midSlot--
        return true
    }
    if carType == 3 && this.smallSlot <= 0 {
        return false
    } else if carType == 3 {
        this.smallSlot--
        return true
    }
    return true    
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */