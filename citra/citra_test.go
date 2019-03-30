package citra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingCitra(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var img = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}

	assert.Equal(t, 3, len(img.Data))
	assert.Equal(t, 0., img.Data[0].Red)
	assert.Equal(t, 25., img.Data[0].Green)
	assert.Equal(t, 50., img.Data[0].Blue)
	assert.Equal(t, 75., img.Data[1].Red)
	assert.Equal(t, 100., img.Data[1].Green)
	assert.Equal(t, 125., img.Data[1].Blue)
	assert.Equal(t, 150., img.Data[2].Red)
	assert.Equal(t, 175., img.Data[2].Green)
	assert.Equal(t, 200., img.Data[2].Blue)
}

func TestConvertCitraToInverse(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var img = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}
	var inv = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}

	inv.Inverse()
	assert.Equal(t, 3, len(inv.Data))

	for a, c := range inv.Data {
		for b, actual := range c.Array().([]float64) {
			expected := 255 - img.Data[a].Array().([]float64)[b]
			assert.Equal(t, expected, actual)
		}
	}
}

func TestConvertCitraToGray(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var gray = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}

	gray.Gray()
	assert.Equal(t, 3, len(gray.Data))

	for _, pixel := range gray.Data {
		pixelArray := pixel.Array().([]float64)
		for j := range pixelArray {
			k := j + 1
			if k == len(pixelArray) {
				k = 0
			}

			assert.Equal(t, pixelArray[j], pixelArray[k])
		}
	}
}

func TestConvertCitraToBinary(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var pixelD = Pixel{Red: 225, Green: 250, Blue: 255}
	var bin = Citra{Data: []Pixel{pixelA, pixelB, pixelC, pixelD}}

	bin.Binary()
	assert.Equal(t, 4, len(bin.Data))

	for i, px := range bin.Data {
		pxArr := px.Array().([]float64)
		for j, e := range pxArr {
			expected := e
			if e > 128 {
				e = 255
			} else {
				e = 0
			}

			actual := bin.Data[i].Array().([]float64)[j]
			assert.Equal(t, expected, actual)
		}
	}
}

func TestCitraToArray(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var before = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}
	var actual = before.Array()
	var expected = [][]float64{{0, 25, 50}, {75, 100, 125}, {150, 175, 200}}
	assert.Equal(t, expected, actual)
}
