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
	data := rand.Perm(100)

	insertionsort(data)

}

func bogosort(data []int) {

	var images []*image.Paletted
	var delays []int

	for true {
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]

		})
		images = append(images, figure.ArrayToScatter(data))
		delays = append(delays, 0)
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

func insertionsort(data []int) {

		var images []*image.Paletted
		var delays []int
		//var disposals []byte

		n := len(data)
		for i := 1; i < n; i++ {
			for j := i; j > 0 && data[j-1] > data[j]; j-- {
				data[j], data[j-1] = data[j-1], data[j]
				if j%5 == 0 {
				images = append(images, figure.ArrayToImage(data))
				delays = append(delays, 0)
				}
					//dis := byte(1)
					//disposals = append(disposals, dis)
			}
		}
		f, err := os.Create("insert.gif")
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)

		gif.EncodeAll(w, &gif.GIF{
			Image: images,
			Delay: delays,
			//Disposal: disposals,
		})
}


func bublesort(data []int) {

	var images []*image.Paletted
	var delays []int
	//var disposals []byte

	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				if j%2 == 0 {
					images = append(images, figure.ArrayToScatter(data))
					delays = append(delays, 1)
				}
				//dis := byte(1)
				//disposals = append(disposals, dis)
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
		//Disposal: disposals,
	})
}
