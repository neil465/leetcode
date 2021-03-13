func judgeCircle(moves string) bool {
	y , x := 0 , 0
	for _, i2 := range moves {
		if string(i2) == "L" {
			x--
		}else if string(i2) == "R" {
			x++
		} else if string(i2) == "U" {
			y++
		} else if string(i2) == "D" {
			y--
		}
	}
	if y == 0 && x == 0 {
		return true
	}
	return false
}