package calendar

import "image"

var (
	up        = decoderDirection(image.Pt(0, -1))
	down      = decoderDirection(image.Pt(0, 1))
	right     = decoderDirection(image.Pt(1, 0))
	left      = decoderDirection(image.Pt(-1, 0))
	rightdown = decoderDirection(image.Pt(1, 1))
	leftup    = decoderDirection(image.Pt(-1, -1))
	rightup   = decoderDirection(image.Pt(1, -1))
	leftdown  = decoderDirection(image.Pt(-1, 1))
)

type decoderDirection image.Point

func (d decoderDirection) Add(d2 decoderDirection) decoderDirection {
	i1 := image.Point(d)
	i2 := image.Point(d2)

	return decoderDirection(i1.Add(i2))
}
