package calendar

import "image"

type Icon string
type DecoderStripeColor string
type DecoderDirection image.Point

type DecoderStripe struct {
	directions []DecoderDirection
	icons      []Icon
}

const (
	IconPlus     Icon = "Plus"
	IconRue      Icon = "Rue"
	IconHexagon  Icon = "Hexagon"
	IconCircle   Icon = "Circle"
	IconStar     Icon = "Star"
	IconMoon     Icon = "Moon"
	IconTriangle Icon = "Triangle"
	IconL        Icon = "L"
	IconY        Icon = "Y"
	IconSquare   Icon = "Square"

	DecoderStripeRed    DecoderStripeColor = "red"
	DecoderStripeBlue   DecoderStripeColor = "blue"
	DecoderStripeYellow DecoderStripeColor = "yellow"
)

var (
	Up        = DecoderDirection(image.Pt(0, -1))
	Down      = DecoderDirection(image.Pt(0, 1))
	Right     = DecoderDirection(image.Pt(1, 0))
	Left      = DecoderDirection(image.Pt(-1, 0))
	RightDown = DecoderDirection(image.Pt(1, 1))
	LeftUp    = DecoderDirection(image.Pt(-1, -1))
	RightUp   = DecoderDirection(image.Pt(1, -1))
	LeftDown  = DecoderDirection(image.Pt(-1, 1))
)

type DecoderBoard struct {
	// The decoderStripes are ordered by the f
	decoderStripes map[DecoderStripeColor]DecoderStripe
}

func NewDecoderBoard() DecoderBoard {
	return DecoderBoard{
		decoderStripes: map[DecoderStripeColor]DecoderStripe{
			DecoderStripeRed: {
				directions: []DecoderDirection{RightDown, LeftUp, RightUp, LeftDown, Up, Down, Right, Left, RightDown, LeftUp},
				icons:      []Icon{IconSquare, IconY, IconL, IconTriangle, IconMoon, IconStar, IconCircle, IconHexagon, IconRue, IconPlus},
			},
			DecoderStripeBlue: {
				directions: []DecoderDirection{Right, Left, RightDown, LeftUp, RightUp, LeftDown, Up, Down, Right, Left},
				icons:      []Icon{IconPlus, IconRue, IconHexagon, IconCircle, IconStar, IconMoon, IconTriangle, IconL, IconY, IconSquare},
			},
			DecoderStripeYellow: {
				directions: []DecoderDirection{Up, Down, Right, Left, RightDown, LeftUp, RightUp, LeftDown, Up, Down},
			},
		},
	}
}
