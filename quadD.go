package piscine

import "fmt"

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < y; i++ {
		line := ""
		for j := 0; j < x; j++ {
			top := i == 0
			bottom := i == y-1
			left := j == 0
			right := j == x-1
			switch {
			case left && (top || bottom):
				line += "A"
			case right && (top || bottom):
				line += "C"
			case top || bottom || left || right:
				line += "B"
			default:
				line += " "
			}
		}
		fmt.Println(line)
	}
}