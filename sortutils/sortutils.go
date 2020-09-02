package sortutils

import "fmt"

//Validate checks whether at dataset is sorted by least to greatest
func Validate(data []int) bool {
	fmt.Println(data)
	n := len(data)
	for i := n - 1; i > 0; i-- {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
