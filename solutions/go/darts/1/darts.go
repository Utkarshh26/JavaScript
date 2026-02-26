package darts

import "math"
func Score(x, y float64) int {
	distance := math.Hypot(x,y)
    if distance > 10 {
        return 0
    }else if distance > 5 && distance <= 10 {
        return 1
    }else if distance > 1 && distance <= 5 {
        return 5
    }else {
        return 10
    }
    return -1
}
