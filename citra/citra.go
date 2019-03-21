package citra

type Citra struct {
	Data []Pixel
}

func (c Citra) toInverse() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.Inverse()
	}

	return n
}

func (c Citra) toGray() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.Inverse()
	}

	return n
}

func (c Citra) toBinary() Citra {
	n := c
	for i, elm := range c.Data {
		n.Data[i] = elm.Binary()
	}

	return n
}
