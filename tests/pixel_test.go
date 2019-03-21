package tests

import (
	"github.com/stretchr/testify/assert"
	"olahcitra/citra"
	"testing"
)

func TestCreatingPixel(t *testing.T) {
	var pixel = citra.Pixel{Red: 100, Green: 200, Blue: 55}
	assert.Equal(t, 100., pixel.Red)
	assert.Equal(t, 200., pixel.Green)
	assert.Equal(t, 55., pixel.Blue)
}

func TestInversePixel(t *testing.T) {
	var before = citra.Pixel{Red: 100, Green: 200, Blue: 55}
	var actual = before.Inverse()
	var expected = citra.Pixel{Red: 155, Green: 55, Blue: 200}

	assert.Equal(t, expected, actual)
}

func TestPixelToGrayScale(t *testing.T) {
	var before = citra.Pixel{Red: 300, Green: 200, Blue: 100}
	var actual = before.GrayScale()

	assert.Equal(t, actual.Red, actual.Green)
	assert.Equal(t, actual.Green, actual.Blue)
	assert.Equal(t, actual.Blue, actual.Red)
}

func TestPixelToBinaryRoundUp(t *testing.T) {
	var before = citra.Pixel{Red: 300, Green: 200, Blue: 100}
	var actual = before.Binary()
	var expected = citra.Pixel{Red: 255, Green: 255, Blue: 255}

	assert.Equal(t, actual.Red, actual.Green)
	assert.Equal(t, actual.Green, actual.Blue)
	assert.Equal(t, actual.Blue, actual.Red)
	assert.Equal(t, expected, actual)
}

func TestPixelToBinaryRoundDown(t *testing.T) {
	var before = citra.Pixel{Red: 100, Green: 25, Blue: 25}
	var actual = before.Binary()
	var expected = citra.Pixel{Red: 0, Green: 0, Blue: 0}

	assert.Equal(t, actual.Red, actual.Green)
	assert.Equal(t, actual.Green, actual.Blue)
	assert.Equal(t, actual.Blue, actual.Red)
	assert.Equal(t, expected, actual)
}
