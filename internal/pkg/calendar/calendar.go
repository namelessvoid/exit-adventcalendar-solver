package calendar

import "image"

type CalendarDoor struct {
	icons []icon
}

type Calendar struct {
	doors map[image.Point]CalendarDoor
}

func NewCalendar() Calendar {
	return Calendar{
		doors: map[image.Point]CalendarDoor{
			image.Pt(0, 0): {icons: []icon{iconTriangle, iconTriangle, iconY}},
			image.Pt(1, 0): {icons: []icon{iconY, iconCircle, iconY}},
			image.Pt(2, 0): {icons: []icon{iconHexagon, iconTriangle, iconMoon}},
			image.Pt(3, 0): {icons: []icon{iconHexagon, iconTriangle, iconPlus}},
			image.Pt(4, 0): {icons: []icon{iconStar, iconStar, iconHexagon}},
			image.Pt(5, 0): {icons: []icon{iconHexagon, iconPlus, iconMoon}},

			image.Pt(0, 1): {icons: []icon{iconCircle, iconHexagon, iconHexagon}},
			image.Pt(1, 1): {icons: []icon{iconHexagon, iconRue, iconTriangle}},
			image.Pt(2, 1): {icons: []icon{iconPlus, iconCircle, iconRue}},
			image.Pt(3, 1): {icons: []icon{iconTriangle, iconY, iconCircle}},
			image.Pt(4, 1): {icons: []icon{iconCircle, iconTriangle, iconCircle}},
			image.Pt(5, 1): {icons: []icon{iconY, iconSquare, iconL}},

			image.Pt(0, 2): {icons: []icon{iconHexagon, iconL, iconTriangle}},
			image.Pt(1, 2): {icons: []icon{iconCircle, iconMoon, iconHexagon}},
			image.Pt(2, 2): {icons: []icon{iconL, iconRue, iconPlus}},
			image.Pt(3, 2): {icons: []icon{iconTriangle, iconHexagon, iconMoon}},
			image.Pt(4, 2): {icons: []icon{iconRue, iconHexagon, iconHexagon}},
			image.Pt(5, 2): {icons: []icon{iconRue, iconRue, iconCircle}},

			image.Pt(0, 3): {icons: []icon{iconCircle, iconL, iconStar}},
			image.Pt(1, 3): {icons: []icon{iconMoon, iconRue, iconRue}},
			image.Pt(2, 3): {icons: []icon{iconHexagon, iconHexagon, iconTriangle}},
			image.Pt(3, 3): {icons: []icon{iconPlus, iconSquare, iconStar}},
			image.Pt(4, 3): {icons: []icon{iconRue, iconPlus, iconRue}},
			image.Pt(5, 3): {icons: []icon{iconL, iconL, iconSquare}},
		},
	}
}

// Returned icons can be handed over to the decoder. The order of the returned
// icons corresponds to the decoder stripes Yellow, Blue, Red.
func (c Calendar) GetIconsOnDoor(coordinates image.Point) []icon {
	return c.doors[coordinates].icons
}

func (c Calendar) GetStart() image.Point {
	return image.Pt(2, 3)
}

func (c Calendar) GetWidth() int {
	return 6
}

func (c Calendar) GetHeight() int {
	return 4
}
