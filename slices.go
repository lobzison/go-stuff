package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    res := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
	    tmp := make([]uint8, dx)
	    for x := 0; x < dx; x++ {
		    tmp[x] = uint8(x)^uint8(y)
		}
		res[y] = tmp
	}
	return res
}

func main() {
	pic.Show(Pic)
}
