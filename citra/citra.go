package citra

type Citra struct {
	Data []Pixel
}

func (c Citra) Inverse() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.Inverse()
	}

	return n
}

func (c Citra) Gray() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.GrayScale()
	}

	return n
}

func (c Citra) Binary() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.Binary()
	}

	return n
}

func (c Citra) Array() interface{} {
	var data [][]float64

	for _, elm := range c.Data {
		data = append(data, elm.Array().([]float64))
	}

	return data
}
