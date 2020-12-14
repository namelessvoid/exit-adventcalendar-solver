package calendar

// The decoderStripeColor defines the color of a stripe.
type decoderStripeColor string

const (
	decoderStripeRed    decoderStripeColor = "red"
	decoderStripeBlue   decoderStripeColor = "blue"
	decoderStripeYellow decoderStripeColor = "yellow"
)

// A decoderStripe represents one of the three paper lines with numbers, icons and directions of
// the cardboard decoder board. Note: We do not need the numbers for the solver.
type decoderStripe struct {
	directions []decoderDirection
	icons      []icon
}

// GetDirectionByIcon returns the arrow on the decoder stripe which corresponds to the
// given icon.
func (d decoderStripe) GetDirectionByIcon(searchedIcon icon) decoderDirection {
	for i, iconIter := range d.icons {
		if iconIter == searchedIcon {
			return d.directions[i]
		}
	}

	panic("Should not happen!")
}
