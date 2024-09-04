package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture{
		picture[i] = make([]uint8 ,dx)
	}
	
	for j := range picture{
		for l := range picture[j]{
			picture[j][l] = uint8(j*l)
		}

	}
	
	
	
	return picture
	
	
}

func main() {
	pic.Show(Pic)
}
