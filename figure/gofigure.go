package figure

import (
	"image"
	"image/color"
	//"fmt"
)

//ArrayToImage plots the array and returns a image
func ArrayToImage(data []int) *image.Paletted {
	//var palette = []color.Color{
	//	color.RGBA{0, 150, 0, 0xff}, color.White, color.Black,
	//}
	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0xff, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0xff, 0xff},
		color.RGBA{0xff, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff},
	}

	width := 1
	spacing := 0

	length := len(data)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{length * (spacing + width), length}

	img := image.NewPaletted(image.Rectangle{upLeft, lowRight}, palette)
	green := color.RGBA{0, 150, 0, 0xff}

	for col := 0; col < length; col++ {
		for height := 0; height <= data[col]; height++ {
			y := length - height
			img.Set(col, y, green)
		}
	}

	return img
}
