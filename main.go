package main

import (
	"fmt"
	"math/rand"

	sutil "./sortutils"
)

func main() {
	data := rand.Perm(100)

	fmt.Println(data)
	//fmt.Println(bogosort(data))
	fmt.Println(bublesort(data))
}

func bogosort(data []int) []int {
	for true {
		rand.Shuffle(len(data), func(i, j int) {
			data[i], data[j] = data[j], data[i]
		})
		if !sutil.Validate(data) {
			continue
		} else {
			break
		}
	}
	return data
}

func bublesort(data []int) []int {
	n := len(data)
	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {

			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				fmt.Println(data)
			}

		}
	}
	return data
}
