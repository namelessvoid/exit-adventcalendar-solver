package calendar

import "image"

type DecoderStripeIcon string
type DecoderStripeColor string
type DecoderDirection image.Point

type DecoderStripe struct {
	directions []DecoderDirection
	icons      []DecoderStripeIcon
}

const (
	DecoderIconPlus     DecoderStripeIcon = "Plus"
	DecoderIconRue      DecoderStripeIcon = "Rue"
	DecoderIconHexagon  DecoderStripeIcon = "Hexagon"
	DecoderIconCircle   DecoderStripeIcon = "Circle"
	DecoderIconStar     DecoderStripeIcon = "Star"
	DecoderIconMoon     DecoderStripeIcon = "Moon"
	DecoderIconTriangle DecoderStripeIcon = "Triangle"
	DecoderIconL        DecoderStripeIcon = "L"
	DecoderIconY        DecoderStripeIcon = "Y"
	DecoderIconSquare   DecoderStripeIcon = "Square"

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
				icons:      []DecoderStripeIcon{DecoderIconSquare, DecoderIconY, DecoderIconL, DecoderIconTriangle, DecoderIconMoon, DecoderIconStar, DecoderIconCircle, DecoderIconHexagon, DecoderIconRue, DecoderIconPlus},
			},
			DecoderStripeBlue: {
				directions: []DecoderDirection{Right, Left, RightDown, LeftUp, RightUp, LeftDown, Up, Down, Right, Left},
				icons:      []DecoderStripeIcon{DecoderIconPlus, DecoderIconRue, DecoderIconHexagon, DecoderIconCircle, DecoderIconStar, DecoderIconMoon, DecoderIconTriangle, DecoderIconL, DecoderIconY, DecoderIconSquare},
			},
			DecoderStripeYellow: {
				directions: []DecoderDirection{Up, Down, Right, Left, RightDown, LeftUp, RightUp, LeftDown, Up, Down},
			},
		},
	}
}
