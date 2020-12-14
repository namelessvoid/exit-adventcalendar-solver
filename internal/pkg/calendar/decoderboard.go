package calendar

type DecoderBoard struct {
	decoderStripes map[decoderStripeColor]decoderStripe
}

func NewDecoderBoard() DecoderBoard {
	return DecoderBoard{
		decoderStripes: map[decoderStripeColor]decoderStripe{
			decoderStripeRed: {
				directions: []decoderDirection{rightdown, leftup, rightup, leftdown, up, down, right, left, rightdown, leftup},
				icons:      []icon{iconSquare, iconY, iconL, iconTriangle, iconMoon, iconStar, iconCircle, iconHexagon, iconRue, iconPlus},
			},
			decoderStripeBlue: {
				directions: []decoderDirection{right, left, rightdown, leftup, rightup, leftdown, up, down, right, left},
				icons:      []icon{iconPlus, iconRue, iconHexagon, iconCircle, iconStar, iconMoon, iconTriangle, iconL, iconY, iconSquare},
			},
			decoderStripeYellow: {
				directions: []decoderDirection{up, down, right, left, rightdown, leftup, rightup, leftdown, up, down},
				icons:      []icon{iconPlus, iconRue, iconHexagon, iconCircle, iconStar, iconMoon, iconTriangle, iconL, iconY, iconSquare},
			},
		},
	}
}

// Icons have to be given in oder Yellow, Blue, Red
func (d DecoderBoard) GetDirectionByIcons(icons []icon) []decoderDirection {
	return []decoderDirection{
		d.decoderStripes[decoderStripeYellow].GetDirectionByIcon(icons[0]),
		d.decoderStripes[decoderStripeBlue].GetDirectionByIcon(icons[1]),
		d.decoderStripes[decoderStripeRed].GetDirectionByIcon(icons[2]),
	}
}
