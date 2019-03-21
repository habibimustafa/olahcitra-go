package main

import "fmt"

func normalizeValue(value int) int {
	if value < 0 {
		value = 0
	}

	if value > 255 {
		value = 255
	}

	return value
}

func toGrayscale(rgb [3]int) [3]int {
	var i int

	for i = 0; i < 3; i++ {
		var gr = int(float32(rgb[0])*0.299) +
			int(float32(rgb[1])*0.587) +
			int(float32(rgb[2])*0.114)

		gr = normalizeValue(gr)
		rgb[i] = gr
	}

	return rgb
}

func main() {
	rgb := [3]int{180, 220, 25}
	fmt.Println(rgb)
	fmt.Println(rgb[0])
	fmt.Println(rgb[1])
	fmt.Println(rgb[2])

	newrgb := toGrayscale(rgb)
	fmt.Println(newrgb)
}
