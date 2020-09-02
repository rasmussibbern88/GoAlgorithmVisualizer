package main

import (
	"bufio"
	"image"
	"image/gif"
	"math/rand"
	"os"

	"./figure"
	sutil "./sortutils"
)

func main() {
	data := rand.Perm(6)

	bogosort(data)

}

func bogosort(data []int) {

	var images []*image.Paletted
	var delays []int

	for true {
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]

		})
		images = append(images, figure.ArrayToImage(data))
		delays = append(delays, 1)
		if !sutil.Validate(data) {
			continue
		} else {
			break
		}
	}
	f, err := os.Create("visbogo.gif")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	gif.EncodeAll(w, &gif.GIF{
		Image:     images,
		Delay:     delays,
		LoopCount: -1,
	})
}

func bublesort(data []int) {

	var images []*image.Paletted
	var delays []int

	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				images = append(images, figure.ArrayToImage(data))
				delays = append(delays, 1)
			}
		}
	}
	f, err := os.Create("vis.gif")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	gif.EncodeAll(w, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
