package solver

import (
	"image"

	"github.com/namelessvoid/exit-solver/internal/pkg/calendar"
)

// A solver is capable of computing the solution of the exit calendar.
type solver struct {
	decoderBoard calendar.DecoderBoard
	calendar     calendar.Calendar
}

// Create a new solver with given decoder board and calendar.
func NewSolver(decoderBoard calendar.DecoderBoard, calendar calendar.Calendar) *solver {
	return &solver{
		decoderBoard: decoderBoard,
		calendar:     calendar,
	}
}

// Solve the calendar.
func (s *solver) Solve() path {
	paths := s.traverse(path{s.calendar.GetStart()})
	return paths
}

// Internally used function to do depth-first traversal.
// If no path is found, it panics because this should never happen.
func (s *solver) traverse(p path) path {
	if len(p) == 13 {
		return p
	}

	for x := 0; x < s.calendar.GetWidth(); x++ {
		for y := 0; y < s.calendar.GetHeight(); y++ {
			tail := p.getTail()

			doorCoordinates := image.Pt(x, y)
			iconsOnDoor := s.calendar.GetIconsOnDoor(doorCoordinates)

			directionsOnDoor := s.decoderBoard.GetDirectionByIcons(iconsOnDoor)
			aggregatedDirection := image.Point(directionsOnDoor[0].Add(directionsOnDoor[1]).Add(directionsOnDoor[2]))

			if doorCoordinates == tail.Add(aggregatedDirection) {
				p = p.appendPoint(doorCoordinates)
				return s.traverse(p)
			}
		}
	}

	panic("No path found!")
}
