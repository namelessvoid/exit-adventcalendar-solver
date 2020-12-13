package calendar

import "image"

type CalendarDoor struct {
	icons []Icon
}

type Calendar struct {
	doors map[image.Point]CalendarDoor
}

func NewCalendar() Calendar {
	return Calendar{
		doors: map[image.Point]CalendarDoor{
			image.Pt(0, 0): {icons: []Icon{IconTriangle, IconTriangle, IconY}},
			image.Pt(1, 0): {icons: []Icon{IconY, IconCircle, IconY}},
			image.Pt(2, 0): {icons: []Icon{IconHexagon, IconTriangle, IconMoon}},
			image.Pt(3, 0): {icons: []Icon{IconHexagon, IconTriangle, IconPlus}},
			image.Pt(4, 0): {icons: []Icon{IconStar, IconStar, IconHexagon}},
			image.Pt(5, 0): {icons: []Icon{IconHexagon, IconPlus, IconMoon}},

			image.Pt(0, 1): {icons: []Icon{IconCircle, IconHexagon, IconHexagon}},
			image.Pt(1, 1): {icons: []Icon{IconHexagon, IconRue, IconTriangle}},
			image.Pt(2, 1): {icons: []Icon{IconPlus, IconCircle, IconRue}},
			image.Pt(3, 1): {icons: []Icon{IconTriangle, IconY, IconCircle}},
			image.Pt(4, 1): {icons: []Icon{IconCircle, IconTriangle, IconCircle}},
			image.Pt(5, 1): {icons: []Icon{IconY, IconSquare, IconL}},

			image.Pt(0, 2): {icons: []Icon{IconHexagon, IconL, IconTriangle}},
			image.Pt(1, 2): {icons: []Icon{IconCircle, IconMoon, IconHexagon}},
			image.Pt(2, 2): {icons: []Icon{IconL, IconRue, IconPlus}},
			image.Pt(3, 2): {icons: []Icon{IconTriangle, IconHexagon, IconMoon}},
			image.Pt(4, 2): {icons: []Icon{IconRue, IconHexagon, IconHexagon}},
			image.Pt(5, 2): {icons: []Icon{IconRue, IconRue, IconCircle}},

			image.Pt(0, 3): {icons: []Icon{IconCircle, IconL, IconStar}},
			image.Pt(1, 3): {icons: []Icon{IconMoon, IconRue, IconRue}},
			image.Pt(2, 3): {icons: []Icon{IconHexagon, IconHexagon, IconTriangle}},
			image.Pt(3, 3): {icons: []Icon{IconPlus, IconSquare, IconStar}},
			image.Pt(4, 3): {icons: []Icon{IconRue, IconPlus, IconRue}},
			image.Pt(5, 3): {icons: []Icon{IconL, IconL, IconSquare}},
		},
	}
}
