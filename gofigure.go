package main

import (
  "image"
  "image/color"
  "image/png"
  "os"
  "math/rand"
  //"fmt"
)

func main() {
  maxval := 100
  size := 100
  width := 2
  spacing := 1

  upLeft := image.Point{0, 0}
  lowRight := image.Point{size*(spacing + width), maxval}

  img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
  green := color.RGBA{0, 150, 0, 0xff}

  n := rand.Perm(size)

  for col := 0; col < len(n); col++ {
    for height := 0; height <= n[col]; height++ {
      y := maxval - height
      img.Set(col,y,green)
    }
  }
  f, _ := os.Create("test2.png")
  png.Encode(f, img)
}



func example() {
  width := 200
  height := 100

  upLeft := image.Point{0, 0}
  lowRight := image.Point{width, height}

  img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

  // Colors are defined by Red, Green, Blue, Alpha uint8 values.
  cyan := color.RGBA{100, 200, 200, 0xff}

  // Set color for each pixel.
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      switch {
        case x < width/2 && y < height/2: // upper left quadrant
        img.Set(x, y, cyan)
        case x >= width/2 && y >= height/2: // lower right quadrant
        img.Set(x, y, color.White)
      default:
        // Use zero value.
      }
    }
  }
  // Encode as PNG.
  f, _ := os.Create("image.png")
  png.Encode(f, img)



}
