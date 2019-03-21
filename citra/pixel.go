package citra

type Pixel struct {
	Red   float64
	Green float64
	Blue  float64
}

func normalize(value float64) float64 {
	if value < 0 {
		value = 0
	}

	if value > 255 {
		value = 255
	}

	return value
}

func (p Pixel) Inverse() Pixel {
	r := normalize(255 - p.Red)
	g := normalize(255 - p.Green)
	b := normalize(255 - p.Blue)
	return Pixel{r, g, b}
}

func (p Pixel) GrayScale() Pixel {
	r := p.Red*0.299 + p.Green*0.587 + p.Blue*0.114
	r = normalize(r)
	return Pixel{r, r, r}
}

func (p Pixel) Binary() Pixel {
	r := (p.Red + p.Green + p.Blue) / 3
	if r <= 128 {
		r = 0
	} else {
		r = 255
	}
	return Pixel{r, r, r}
}
