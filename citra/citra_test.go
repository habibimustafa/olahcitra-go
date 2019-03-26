package citra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatingCitra(t *testing.T) {
	var pixelA = Pixel{Red: 0, Green: 25, Blue: 50}
	var pixelB = Pixel{Red: 75, Green: 100, Blue: 125}
	var pixelC = Pixel{Red: 150, Green: 175, Blue: 200}
	var gambar = Citra{Data: []Pixel{pixelA, pixelB, pixelC}}

	assert.Equal(t, 3, len(gambar.Data))
	assert.Equal(t, 0., gambar.Data[0].Red)
	assert.Equal(t, 25., gambar.Data[0].Green)
	assert.Equal(t, 50., gambar.Data[0].Blue)
	assert.Equal(t, 75., gambar.Data[1].Red)
	assert.Equal(t, 100., gambar.Data[1].Green)
	assert.Equal(t, 125., gambar.Data[1].Blue)
	assert.Equal(t, 150., gambar.Data[2].Red)
	assert.Equal(t, 175., gambar.Data[2].Green)
	assert.Equal(t, 200., gambar.Data[2].Blue)
}